/**
 * Copyright 2002-2012 Evgeny Gryaznov
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package org.textmapper.lapg.parser;

import org.textmapper.lapg.LapgCore;
import org.textmapper.lapg.api.*;
import org.textmapper.lapg.api.builder.GrammarBuilder;
import org.textmapper.lapg.api.builder.RuleBuilder;
import org.textmapper.lapg.api.regex.RegexContext;
import org.textmapper.lapg.api.regex.RegexMatcher;
import org.textmapper.lapg.api.regex.RegexParseException;
import org.textmapper.lapg.api.regex.RegexPart;
import org.textmapper.lapg.common.FormatUtil;
import org.textmapper.lapg.gen.TemplateStaticMethods;
import org.textmapper.lapg.parser.LapgLexer.ErrorReporter;
import org.textmapper.lapg.parser.LapgLexer.LapgSymbol;
import org.textmapper.lapg.parser.LapgLexer.Lexems;
import org.textmapper.lapg.parser.LapgTree.LapgProblem;
import org.textmapper.lapg.parser.LapgTree.TextSource;
import org.textmapper.lapg.parser.ast.*;
import org.textmapper.templates.api.types.IClass;
import org.textmapper.templates.api.types.IFeature;
import org.textmapper.templates.api.types.IType;
import org.textmapper.templates.types.TiExpressionBuilder;
import org.textmapper.templates.types.TypesRegistry;
import org.textmapper.templates.types.TypesUtil;

import java.io.IOException;
import java.util.*;

public class LapgResolver {

	public static final String RESOLVER_SOURCE = "problem.resolver"; //$NON-NLS-1$

	private final LapgTree<AstRoot> tree;
	private final TypesRegistry types;
	private String myTypesPackage;

	private final Map<String, Symbol> symbolsMap = new HashMap<String, Symbol>();
	private final Map<Symbol, String> identifierMap = new HashMap<Symbol, String>();
	private final Map<SourceElement, Map<String, Object>> annotationsMap = new HashMap<SourceElement, Map<String, Object>>();
	private final Map<SourceElement, TextSourceElement> codeMap = new HashMap<SourceElement, TextSourceElement>();

	private Map<String, Object> options;
	private GrammarBuilder builder;
	private boolean hasInputs = false;

	public LapgResolver(LapgTree<AstRoot> tree, TypesRegistry types) {
		this.tree = tree;
		this.types = types;
	}

	public LapgGrammar resolve() {
		if (tree.getRoot() == null) {
			return null;
		}
		myTypesPackage = getTypesPackage();

		builder = LapgCore.createBuilder();
		symbolsMap.put("eoi", builder.getEoi());
		collectLexems();

		if (tree.getRoot().getGrammar() != null) {
			collectNonTerminals();
			collectRules();
			collectDirectives();

			if (!hasInputs) {
				Symbol input = symbolsMap.get("input");
				if (input == null) {
					error(tree.getRoot(), "no input non-terminal");
				} else if (input.isTerm()) {
					error(tree.getRoot(), "input should be non-terminal");
				} else {
					builder.addInput(input, true, input);
				}
			}

			// TODO collect layout symbols
		}

		collectOptions();
		TextSourceElement templates = getTemplates();
		String copyrightHeader = extractCopyright();

		Grammar g = builder.create();

		UniqueNameHelper helper = new UniqueNameHelper();
		for (Symbol s : g.getSymbols()) {
			String name = s.getName();
			if (FormatUtil.isIdentifier(name)) {
				helper.markUsed(name);
			}
		}
		for (int i = 0; i < g.getSymbols().length; i++) {
			Symbol sym = g.getSymbols()[i];
			identifierMap.put(sym, helper.generateId(sym.getName(), i));
		}

		return new LapgGrammar(g, templates, !tree.getErrors().isEmpty(), options, copyrightHeader,
				identifierMap, annotationsMap, codeMap);
	}

	private TextSourceElement getTemplates() {
		int offset = tree.getRoot() != null ? tree.getRoot().getTemplatesStart() : -1;
		char[] text = tree.getSource().getContents();
		if (offset < text.length && offset != -1) {
			return new AstNode(tree.getSource(), offset, text.length) {
				@Override
				public void accept(AbstractVisitor v) {
				}
			};
		}
		return null;
	}

	private Symbol create(AstIdentifier id, String type, int kind, Symbol softClass) {
		String name = id.getName();
		if (symbolsMap.containsKey(name)) {
			Symbol sym = symbolsMap.get(name);
			if (sym.getKind() != kind) {
				error(id, "redeclaration of " + sym.kindAsString() + ": " + name);
			} else if (!UniqueNameHelper.safeEquals(sym.getType(), type) && !(kind == Symbol.KIND_SOFTTERM && type == null)) {
				error(id,
						"redeclaration of type: " + (type == null ? "<empty>" : type) + " instead of "
								+ (sym.getType() == null ? "<empty>" : sym.getType()));
			} else if (kind == Symbol.KIND_SOFTTERM && softClass != sym.getSoftClass()) {
				error(id, "redeclaration of soft class: " + (softClass == null ? "<undefined>" : softClass.getName())
						+ " instead of " + (sym.getSoftClass() == null ? "<undefined>" : sym.getSoftClass().getName()));
			}
			return sym;
		} else {
			Symbol sym = kind == Symbol.KIND_SOFTTERM
					? builder.addSoftSymbol(name, softClass, id)
					: builder.addSymbol(kind, name, type, id);
			symbolsMap.put(name, sym);
			return sym;
		}
	}

	private Map<Symbol, Integer> lastIndex = new HashMap<Symbol, Integer>();

	private Symbol createNested(int kind, String type, Symbol outer, Symbol softClass, IAstNode source) {
		int index = lastIndex.containsKey(outer) ? lastIndex.get(outer) : 1;
		String base_ = outer.getName() + "$";
		while (symbolsMap.containsKey(base_ + index)) {
			index++;
		}
		String name = base_ + index;
		Symbol sym = kind == Symbol.KIND_SOFTTERM
				? builder.addSoftSymbol(name, softClass, source)
				: builder.addSymbol(kind, name, type, source);
		symbolsMap.put(name, sym);
		lastIndex.put(outer, index + 1);
		return sym;
	}

	private Symbol resolve(AstReference id) {
		String name = id.getName();
		Symbol sym = symbolsMap.get(name);
		if (sym == null) {
			if (name.length() > 3 && name.endsWith("opt")) {
				sym = symbolsMap.get(name.substring(0, name.length() - 3));
				if (sym != null) {
					Symbol symopt = create(
							new AstIdentifier(id.getName(), id.getInput(), id.getOffset(), id.getEndOffset()),
							sym.getType(), Symbol.KIND_NONTERM, null);
					RuleBuilder rb = builder.rule(null, symopt, id);
					rb.create();
					rb.addPart(null, sym, null, id);
					rb.create();
					return symopt;
				}
			}
			error(id, name + " cannot be resolved");
		}
		return sym;
	}

	private int convert(AstGroupsSelector selector) {
		int result = 0;
		for (Integer i : selector.getGroups()) {
			if (i == null || i < 0 || i > 31) {
				error(selector, "group id should be in range 0..31");
				return 1;
			} else if ((result & 1 << i) != 0) {
				error(selector, "duplicate group id: " + i);
				return 1;
			} else {
				result |= 1 << i;
			}
		}
		if (result == 0) {
			error(selector, "empty group set");
			return 1;
		}
		return result;
	}

	private Lexem getClassLexem(Map<Lexem, RegexMatcher> classMatchers, AstLexeme l, RegexPart regex) {
		Lexem result = null;
		AstLexemAttrs attrs = l.getAttrs();
		int kind = attrs == null ? Lexem.KIND_NONE : attrs.getKind();
		if (regex.isConstant() && kind != Lexem.KIND_CLASS) {
			for (Lexem classLexem : classMatchers.keySet()) {
				if (l.getGroups() != classLexem.getGroups()) {
					continue;
				}
				RegexMatcher m = classMatchers.get(classLexem);
				if (m.matches(regex.getConstantValue())) {
					if (result != null) {
						error(l, "lexem matches two classes `" + result.getSymbol().getName() + "' and `"
								+ classLexem.getSymbol().getName() + "', using first");
					} else {
						result = classLexem;
					}
				}
			}
		}
		return result;
	}

	private void collectLexems() {
		List<Lexem> classLexems = new LinkedList<Lexem>();
		Map<String, RegexPart> namedPatternsMap = new HashMap<String, RegexPart>();

		// Step 1. Process class lexems, named patterns & groups.

		int groups = 1;
		for (AstLexerPart clause : tree.getRoot().getLexer()) {
			if (clause instanceof AstLexeme) {
				AstLexeme lexeme = (AstLexeme) clause;
				lexeme.setGroups(groups);
				AstLexemAttrs attrs = lexeme.getAttrs();
				if (attrs == null || attrs.getKind() != Lexem.KIND_CLASS) {
					continue;
				}
				if (lexeme.getRegexp() == null) {
					error(lexeme, "class lexem without regular expression, ignored");
					continue;
				}

				Symbol s = create(lexeme.getName(), lexeme.getType(), Symbol.KIND_TERM, null);
				RegexPart regex;
				try {
					regex = LapgCore.parse(s.getName(), lexeme.getRegexp().getRegexp());
				} catch (RegexParseException e) {
					error(lexeme.getRegexp(), e.getMessage());
					continue;
				}

				Lexem liLexem = builder.addLexem(Lexem.KIND_CLASS, s, regex, lexeme.getGroups(), lexeme.getPriority(),
						null, lexeme);
				classLexems.add(liLexem);
				codeMap.put(liLexem, lexeme.getCode());
			} else if (clause instanceof AstGroupsSelector) {
				groups = convert((AstGroupsSelector) clause);
			} else if (clause instanceof AstNamedPattern) {
				AstNamedPattern astpattern = (AstNamedPattern) clause;
				String name = astpattern.getName();
				RegexPart regex;
				try {
					regex = LapgCore.parse(name, astpattern.getRegexp().getRegexp());
				} catch (RegexParseException e) {
					error(astpattern.getRegexp(), e.getMessage());
					continue;
				}
				if (namedPatternsMap.get(name) != null) {
					error(astpattern, "redeclaration of named pattern `" + name + "'");
				} else {
					builder.addPattern(name, regex, astpattern);
					namedPatternsMap.put(name, regex);
				}
			}
		}

		// Step 2. Process other lexems. Match soft lexems with their classes.

		RegexContext context = LapgCore.createContext(namedPatternsMap);
		Map<Lexem, RegexMatcher> classMatchers = new LinkedHashMap<Lexem, RegexMatcher>();
		for (Lexem clLexem : classLexems) {
			classMatchers.put(clLexem, LapgCore.createMatcher(clLexem.getRegexp(), context));
		}

		for (AstLexerPart clause : tree.getRoot().getLexer()) {
			if (clause instanceof AstLexeme) {
				AstLexeme lexeme = (AstLexeme) clause;
				AstLexemAttrs attrs = lexeme.getAttrs();
				int kind = attrs == null ? Lexem.KIND_NONE : attrs.getKind();
				if (kind == Lexem.KIND_CLASS) {
					continue;
				}
				if (lexeme.getRegexp() != null) {
					String name = lexeme.getName().getName();
					RegexPart regex;
					try {
						regex = LapgCore.parse(name, lexeme.getRegexp().getRegexp());
					} catch (RegexParseException e) {
						error(lexeme.getRegexp(), e.getMessage());
						continue;
					}

					if (kind == Lexem.KIND_SOFT && lexeme.getCode() != null) {
						error(lexeme.getCode(), "soft lexem `" + lexeme.getName().getName()
								+ "' cannot have a semantic action");
					}
					Lexem classLexem = getClassLexem(classMatchers, lexeme, regex);
					Symbol softClass = null;
					if (kind == Lexem.KIND_SOFT) {
						if (classLexem == null) {
							if (!regex.isConstant()) {
								error(lexeme, "soft lexem `" + name + "' should have a constant regexp");
							} else {
								error(lexeme, "soft lexem `" + name + "' doesn't match any class lexem");
							}
							kind = Lexem.KIND_NONE;
						} else {
							softClass = classLexem.getSymbol();

							String type = lexeme.getType();
							String classtype = softClass.getType();
							if (type != null && !type.equals(classtype)) {
								if (classtype == null) {
									classtype = "<no type>";
								}
								error(lexeme, "soft terminal `" + name + "' overrides base type: expected `"
										+ classtype + "', found `" + type + "'");
							}
						}
					}

					Symbol s = create(lexeme.getName(), lexeme.getType(),
							kind == Lexem.KIND_SOFT ? Symbol.KIND_SOFTTERM : Symbol.KIND_TERM, softClass);
					Lexem liLexem = builder.addLexem(kind, s, regex, lexeme.getGroups(), lexeme.getPriority(),
							classLexem, lexeme);
					codeMap.put(liLexem, lexeme.getCode());
				} else {
					if (kind == Lexem.KIND_SOFT) {
						error(lexeme, "soft lexem `" + lexeme.getName().getName() + "' should have regular expression");
					}
					create(lexeme.getName(), lexeme.getType(), Symbol.KIND_TERM, null);
				}
			}
		}
	}

	private void addSymbolAnnotations(AstIdentifier id, Map<String, Object> annotations) {
		if (annotations != null) {
			Symbol sym = symbolsMap.get(id.getName());
			Map<String, Object> symAnnotations = annotationsMap.get(sym);
			if (symAnnotations == null) {
				symAnnotations = new HashMap<String, Object>();
				annotationsMap.put(sym, symAnnotations);
			}
			for (Map.Entry<String, Object> ann : annotations.entrySet()) {
				if (symAnnotations.containsKey(ann.getKey())) {
					error(id, "redeclaration of annotation `" + ann.getKey() + "' for non-terminal: " + id.getName()
							+ ", skipped");
				} else {
					symAnnotations.put(ann.getKey(), ann.getValue());
				}
			}
		}
	}

	private void collectNonTerminals() {
		for (AstGrammarPart clause : tree.getRoot().getGrammar()) {
			if (clause instanceof AstNonTerm) {
				AstNonTerm nonterm = (AstNonTerm) clause;
				create(nonterm.getName(), nonterm.getType(), Symbol.KIND_NONTERM, null);
			}
		}
		for (AstGrammarPart clause : tree.getRoot().getGrammar()) {
			if (clause instanceof AstNonTerm) {
				AstNonTerm nonterm = (AstNonTerm) clause;
				addSymbolAnnotations(nonterm.getName(), convert(nonterm.getAnnotations(), "AnnotateSymbol"));
			}
		}
	}

	private void createRule(Symbol left, AstRule right) {
		RuleBuilder rule = builder.rule(right.getAlias(), left, right);
		List<AstRulePart> list = right.getList();
		AstCode lastAction = null;
		if (list != null) {
			AstRulePart last = list.size() > 0 ? list.get(list.size() - 1) : null;
			if (last instanceof AstCode) {
				lastAction = (AstCode) last;
			}

			for (AstRulePart part : list) {
				if (part instanceof AstCode) {
					AstCode astCode = (AstCode) part;
					if (astCode != lastAction) {
						Symbol codeSym = createNested(Symbol.KIND_NONTERM, null, left, null, astCode);
						Rule actionRule = builder.rule(null, codeSym, astCode).create();
						codeMap.put(actionRule, astCode);
						rule.addPart(null, codeSym, null, astCode);
					}

				} else if (part instanceof AstRuleSymbol) {
					AstRuleSymbol rs = (AstRuleSymbol) part;
					Symbol sym = resolve(rs.getSymbol());
					if (sym != null) {
						// TODO check duplicate alias
						SymbolRef ref = rule.addPart(rs.getAlias(), sym, convertLA(rs.getAnnotations()), rs.getSymbol());
						annotationsMap.put(ref, convert(rs.getAnnotations(), "AnnotateReference"));
					}
				}
			}
		}
		AstRuleAttribute ruleAttribute = right.getAttribute();
		AstReference rulePrio = ruleAttribute instanceof AstPrioClause ? ((AstPrioClause) ruleAttribute).getReference()
				: null;
		if (rulePrio != null) {
			Symbol prio = resolve(rulePrio);
			if (prio != null) {
				rule.setPriority(prio);
			}
		}

		// TODO store %shift attribute
		// TODO check prio is term
		// TODO check right.getAnnotations().getNegativeLA() == null
		Rule result = rule.create();
		annotationsMap.put(result, convert(right.getAnnotations(), "AnnotateRule"));
		codeMap.put(result, lastAction);
	}

	private void collectRules() {
		for (AstGrammarPart clause : tree.getRoot().getGrammar()) {
			if (clause instanceof AstNonTerm) {
				AstNonTerm nonterm = (AstNonTerm) clause;
				Symbol left = symbolsMap.get(nonterm.getName().getName());
				if (left == null || left.getKind() != Symbol.KIND_NONTERM) {
					continue; /* error is already reported */
				}
				for (AstRule right : nonterm.getRules()) {
					if (!right.hasSyntaxError()) {
						createRule(left, right);
					}
				}
			}
		}
	}

	private List<Symbol> resolve(List<AstReference> input) {
		List<Symbol> result = new ArrayList<Symbol>(input.size());
		for (AstReference id : input) {
			Symbol sym = resolve(id);
			if (sym != null) {
				result.add(sym);
			}
		}
		return result;
	}

	private void collectDirectives() {
		for (AstGrammarPart clause : tree.getRoot().getGrammar()) {
			if (clause instanceof AstDirective) {
				AstDirective directive = (AstDirective) clause;
				String key = directive.getKey();
				List<Symbol> val = resolve(directive.getSymbols());
				int prio;
				if (key.equals("left")) {
					prio = Prio.LEFT;
				} else if (key.equals("right")) {
					prio = Prio.RIGHT;
				} else if (key.equals("nonassoc")) {
					prio = Prio.NONASSOC;
				} else {
					error(directive, "unknown directive identifier used: `" + key + "`");
					continue;
				}
				builder.addPrio(prio, val, directive);
			} else if (clause instanceof AstInputDirective) {
				List<AstInputRef> refs = ((AstInputDirective) clause).getInputRefs();
				for (AstInputRef inputRef : refs) {
					Symbol sym = resolve(inputRef.getReference());
					boolean hasEoi = !inputRef.isNonEoi();
					if (sym != null) {
						builder.addInput(sym, hasEoi, inputRef);
						hasInputs = true;
					}
				}
			}
		}
	}

	private String getTypesPackage() {
		if (tree.getRoot().getOptions() != null) {
			for (AstOptionPart option : tree.getRoot().getOptions()) {
				if (option instanceof AstOption && ((AstOption) option).getKey().equals("lang")) {
					AstExpression expression = ((AstOption) option).getValue();
					if (expression instanceof AstLiteralExpression) {
						return ((AstLiteralExpression) expression).getLiteral().toString();
					}
				}
			}
		}

		return "common";
	}

	private void collectOptions() {
		options = new HashMap<String, Object>();

		// Load class
		IClass optionsClass = types.getClass(myTypesPackage + ".Options", null);
		if (optionsClass == null) {
			error(tree.getRoot(), "cannot load options class `" + myTypesPackage + ".Options`");
			return;
		}

		// fill default values
		for (IFeature feature : optionsClass.getFeatures()) {
			Object value = feature.getDefaultValue();
			if (value != null) {
				options.put(feature.getName(), value);
			}
		}

		// overrides
		if (tree.getRoot().getOptions() == null) {
			return;
		}
		for (AstOptionPart option : tree.getRoot().getOptions()) {
			if (option instanceof AstOption) {
				String key = ((AstOption) option).getKey();
				IFeature feature = optionsClass.getFeature(key);
				if (feature == null) {
					error(option, "unknown option `" + key + "`");
					continue;
				}

				AstExpression value = ((AstOption) option).getValue();
				options.put(key, convertExpression(value, feature.getType()));
			}
		}
	}

	private void error(IAstNode n, String message) {
		tree.getErrors().add(new LapgResolverProblem(LapgTree.KIND_ERROR, n.getOffset(), n.getEndOffset(), message));
	}

	private Collection<Symbol> convertLA(AstRuleAnnotations astAnnotations) {
		if (astAnnotations == null || astAnnotations.getNegativeLA() == null) {
			return null;
		}

		List<AstReference> unwantedSymbols = astAnnotations.getNegativeLA().getUnwantedSymbols();
		List<Symbol> resolved = resolve(unwantedSymbols);
		if (resolved.size() == 0) {
			return null;
		}

		return resolved;
	}

	@SuppressWarnings("unchecked")
	private Map<String, Object> convert(AstAnnotations astAnnotations, String kind) {
		if (astAnnotations == null || astAnnotations.getAnnotations() == null) {
			return null;
		}

		// Load class
		IClass annoClass = types.getClass(myTypesPackage + "." + kind, null);
		if (annoClass == null) {
			error(astAnnotations, "cannot load class `" + myTypesPackage + "." + kind + "`");
			return null;
		}

		List<AstNamedEntry> list = astAnnotations.getAnnotations();
		Map<String, Object> result = new HashMap<String, Object>();
		for (AstNamedEntry entry : list) {
			if (entry.hasSyntaxError()) {
				continue;
			}
			String name = entry.getName();
			IFeature feature = annoClass.getFeature(name);
			if (feature == null) {
				error(entry, "unknown annotation `" + name + "`");
				continue;
			}

			IType expected = feature.getType();

			AstExpression expr = entry.getExpression();
			if (expr == null) {
				if (!TypesUtil.isBooleanType(expected)) {
					error(entry, "expected value of type `" + expected.toString() + "` instead of boolean");
					continue;
				}
				result.put(name, Boolean.TRUE);
			} else {
				result.put(name, convertExpression(expr, expected));
			}
		}
		return result;
	}

	@SuppressWarnings("unchecked")
	private Object convertExpression(AstExpression expression, IType type) {
		return new TiExpressionBuilder<AstExpression>() {
			@Override
			public IClass resolveType(String className) {
				return types.getClass(className, null);
			}

			@Override
			public Object resolve(AstExpression expression, IType type) {
				if (expression instanceof AstInstance) {
					List<AstNamedEntry> list = ((AstInstance) expression).getEntries();
					Map<String, AstExpression> props = new HashMap<String, AstExpression>();
					if (list != null) {
						for (AstNamedEntry entry : list) {
							if (entry.hasSyntaxError()) {
								continue;
							}
							props.put(entry.getName(), entry.getExpression());
						}
					}
					String name = ((AstInstance) expression).getClassName().getName();
					if (name.indexOf('.') < 0) {
						name = myTypesPackage + "." + name;
					}
					return convertNew(expression, name, props, type);
				}
				if (expression instanceof AstArray) {
					List<AstExpression> list = ((AstArray) expression).getExpressions();
					return convertArray(expression, list, type);
				}
				if (expression instanceof AstReference) {
					IClass symbolClass = types.getClass("common.Symbol", null);
					if (symbolClass == null) {
						report(expression, "cannot load class `common.Symbol`");
						return null;
					}
					if (!symbolClass.isSubtypeOf(type)) {
						report(expression, "`" + symbolClass.toString() + "` is not a subtype of `" + type.toString()
								+ "`");
						return null;
					}
					return LapgResolver.this.resolve((AstReference) expression);
				}
				if (expression instanceof AstLiteralExpression) {
					Object literal = ((AstLiteralExpression) expression).getLiteral();
					return convertLiteral(expression, literal, type);
				}
				return null;
			}

			@Override
			public void report(AstExpression expression, String message) {
				error(expression, message);
			}
		}.resolve(expression, type);
	}

	private String extractCopyright() {
		TextSource source = tree.getSource();
		final boolean[] hasErrors = new boolean[]{false};
		ErrorReporter reporter = new ErrorReporter() {
			@Override
			public void error(int start, int end, int line, String s) {
				hasErrors[0] = true;
			}
		};

		try {
			LapgLexer lexer = new LapgLexer(source.getStream(), reporter);
			lexer.setSkipComments(false);
			List<String> headers = new LinkedList<String>();

			LapgSymbol sym = lexer.next();
			int lastline = 0;
			StringBuilder sb = new StringBuilder();
			while (sym.lexem == Lexems._skip_comment && source.columnForOffset(sym.offset) == 0) {
				String val = lexer.current().substring(1);
				if (sym.line > lastline + 1 && sb.length() > 0) {
					headers.add(sb.toString());
					sb.setLength(0);
				}
				lastline = sym.line;
				if (!(sym.line == 1 && val.startsWith("!"))) {
					sb.append(val).append('\n');
				}
				sym = lexer.next();
			}
			if (hasErrors[0]) {
				return null;
			}
			if (sb.length() > 0) {
				headers.add(sb.toString());
			}
			for (String s : headers) {
				if (s.toLowerCase().contains("license")) {
					return new TemplateStaticMethods().shiftLeft(s);
				}
			}

		} catch (IOException e) {
			/* ignore */
		}

		return null;
	}

	private static class LapgResolverProblem extends LapgProblem {
		private static final long serialVersionUID = 3810706800688899470L;

		public LapgResolverProblem(int kind, int offset, int endoffset, String message) {
			super(kind, offset, endoffset, message, null);
		}

		@Override
		public String getSource() {
			return RESOLVER_SOURCE;
		}
	}
}
