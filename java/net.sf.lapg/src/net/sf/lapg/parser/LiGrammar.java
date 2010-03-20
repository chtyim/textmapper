package net.sf.lapg.parser;

import java.util.Map;

import net.sf.lapg.api.Grammar;
import net.sf.lapg.api.Lexem;
import net.sf.lapg.api.Prio;
import net.sf.lapg.api.Rule;
import net.sf.lapg.api.Symbol;

public class LiGrammar implements Grammar {

	private final Symbol[] symbols;
	private final Rule[] rules;
	private final Prio[] priorities;
	private final Lexem[] lexems;

	private final Symbol[] inputs;
	private final Symbol eoi;
	private final Symbol error;

	private final Map<String, Object> options;
	private final String templates;
	private final int terminals;
	private final boolean hasErrors;

	public LiGrammar(Symbol[] symbols, Rule[] rules, Prio[] priorities, Lexem[] lexems, Symbol[] inputs, Symbol eoi,
			Symbol error, Map<String, Object> options, String templates, int terminals, boolean hasErrors) {
		this.symbols = symbols;
		this.rules = rules;
		this.priorities = priorities;
		this.lexems = lexems;
		this.inputs = inputs;
		this.eoi = eoi;
		this.error = error;
		this.options = options;
		this.templates = templates;
		this.terminals = terminals;
		this.hasErrors = hasErrors;
	}

	public Symbol[] getSymbols() {
		return symbols;
	}

	public Rule[] getRules() {
		return rules;
	}

	public Prio[] getPriorities() {
		return priorities;
	}

	public Lexem[] getLexems() {
		return lexems;
	}

	public Symbol[] getInput() {
		return inputs;
	}

	public Symbol getEoi() {
		return eoi;
	}

	public Symbol getError() {
		return error;
	}

	public Map<String, Object> getOptions() {
		return options;
	}

	public String getTemplates() {
		return templates;
	}

	public int getTerminals() {
		return terminals;
	}

	public boolean hasErrors() {
		return hasErrors;
	}

	public boolean hasActions() {
		for (Rule r : rules) {
			if (r.getAction() != null) {
				return true;
			}
		}
		return false;
	}

	public boolean hasLexemActions() {
		for (Lexem l : lexems) {
			if (l.getAction() != null) {
				return true;
			}
		}
		return false;
	}
}
