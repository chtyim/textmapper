// generated by Textmapper; DO NOT EDIT

package selector

import (
	"github.com/inspirer/textmapper/tm-parsers/js"
)

type Selector func(nt js.NodeType) bool

var (
	Any                        = func(t js.NodeType) bool { return true }
	Abstract                   = func(t js.NodeType) bool { return t == js.Abstract }
	AccessibilityModifier      = func(t js.NodeType) bool { return t == js.AccessibilityModifier }
	AdditiveExpression         = func(t js.NodeType) bool { return t == js.AdditiveExpression }
	Arguments                  = func(t js.NodeType) bool { return t == js.Arguments }
	ArrayLiteral               = func(t js.NodeType) bool { return t == js.ArrayLiteral }
	ArrayPattern               = func(t js.NodeType) bool { return t == js.ArrayPattern }
	ArrayType                  = func(t js.NodeType) bool { return t == js.ArrayType }
	ArrowFunction              = func(t js.NodeType) bool { return t == js.ArrowFunction }
	AssignmentExpression       = func(t js.NodeType) bool { return t == js.AssignmentExpression }
	AssignmentOperator         = func(t js.NodeType) bool { return t == js.AssignmentOperator }
	AsyncArrowFunction         = func(t js.NodeType) bool { return t == js.AsyncArrowFunction }
	AsyncFunction              = func(t js.NodeType) bool { return t == js.AsyncFunction }
	AsyncFunctionExpression    = func(t js.NodeType) bool { return t == js.AsyncFunctionExpression }
	AsyncMethod                = func(t js.NodeType) bool { return t == js.AsyncMethod }
	AwaitExpression            = func(t js.NodeType) bool { return t == js.AwaitExpression }
	BindingIdentifier          = func(t js.NodeType) bool { return t == js.BindingIdentifier }
	BindingRestElement         = func(t js.NodeType) bool { return t == js.BindingRestElement }
	BitwiseANDExpression       = func(t js.NodeType) bool { return t == js.BitwiseANDExpression }
	BitwiseORExpression        = func(t js.NodeType) bool { return t == js.BitwiseORExpression }
	BitwiseXORExpression       = func(t js.NodeType) bool { return t == js.BitwiseXORExpression }
	Block                      = func(t js.NodeType) bool { return t == js.Block }
	Body                       = func(t js.NodeType) bool { return t == js.Body }
	BreakStatement             = func(t js.NodeType) bool { return t == js.BreakStatement }
	CallExpression             = func(t js.NodeType) bool { return t == js.CallExpression }
	CallSignature              = func(t js.NodeType) bool { return t == js.CallSignature }
	Case                       = func(t js.NodeType) bool { return t == js.Case }
	Catch                      = func(t js.NodeType) bool { return t == js.Catch }
	Class                      = func(t js.NodeType) bool { return t == js.Class }
	ClassBody                  = func(t js.NodeType) bool { return t == js.ClassBody }
	ClassExpr                  = func(t js.NodeType) bool { return t == js.ClassExpr }
	CommaExpression            = func(t js.NodeType) bool { return t == js.CommaExpression }
	ComputedPropertyName       = func(t js.NodeType) bool { return t == js.ComputedPropertyName }
	ConciseBody                = func(t js.NodeType) bool { return t == js.ConciseBody }
	ConditionalExpression      = func(t js.NodeType) bool { return t == js.ConditionalExpression }
	ConstructSignature         = func(t js.NodeType) bool { return t == js.ConstructSignature }
	ConstructorType            = func(t js.NodeType) bool { return t == js.ConstructorType }
	ContinueStatement          = func(t js.NodeType) bool { return t == js.ContinueStatement }
	DebuggerStatement          = func(t js.NodeType) bool { return t == js.DebuggerStatement }
	Default                    = func(t js.NodeType) bool { return t == js.Default }
	DefaultParameter           = func(t js.NodeType) bool { return t == js.DefaultParameter }
	DoWhileStatement           = func(t js.NodeType) bool { return t == js.DoWhileStatement }
	ElementBinding             = func(t js.NodeType) bool { return t == js.ElementBinding }
	EmptyDecl                  = func(t js.NodeType) bool { return t == js.EmptyDecl }
	EmptyStatement             = func(t js.NodeType) bool { return t == js.EmptyStatement }
	EqualityExpression         = func(t js.NodeType) bool { return t == js.EqualityExpression }
	ExponentiationExpression   = func(t js.NodeType) bool { return t == js.ExponentiationExpression }
	ExportClause               = func(t js.NodeType) bool { return t == js.ExportClause }
	ExportDeclaration          = func(t js.NodeType) bool { return t == js.ExportDeclaration }
	ExportDefault              = func(t js.NodeType) bool { return t == js.ExportDefault }
	ExportSpecifier            = func(t js.NodeType) bool { return t == js.ExportSpecifier }
	ExpressionStatement        = func(t js.NodeType) bool { return t == js.ExpressionStatement }
	Extends                    = func(t js.NodeType) bool { return t == js.Extends }
	Finally                    = func(t js.NodeType) bool { return t == js.Finally }
	ForBinding                 = func(t js.NodeType) bool { return t == js.ForBinding }
	ForCondition               = func(t js.NodeType) bool { return t == js.ForCondition }
	ForFinalExpression         = func(t js.NodeType) bool { return t == js.ForFinalExpression }
	ForInStatement             = func(t js.NodeType) bool { return t == js.ForInStatement }
	ForInStatementWithVar      = func(t js.NodeType) bool { return t == js.ForInStatementWithVar }
	ForOfStatement             = func(t js.NodeType) bool { return t == js.ForOfStatement }
	ForOfStatementWithVar      = func(t js.NodeType) bool { return t == js.ForOfStatementWithVar }
	ForStatement               = func(t js.NodeType) bool { return t == js.ForStatement }
	ForStatementWithVar        = func(t js.NodeType) bool { return t == js.ForStatementWithVar }
	Function                   = func(t js.NodeType) bool { return t == js.Function }
	FunctionExpression         = func(t js.NodeType) bool { return t == js.FunctionExpression }
	FunctionType               = func(t js.NodeType) bool { return t == js.FunctionType }
	Generator                  = func(t js.NodeType) bool { return t == js.Generator }
	GeneratorExpression        = func(t js.NodeType) bool { return t == js.GeneratorExpression }
	GeneratorMethod            = func(t js.NodeType) bool { return t == js.GeneratorMethod }
	Getter                     = func(t js.NodeType) bool { return t == js.Getter }
	IdentifierReference        = func(t js.NodeType) bool { return t == js.IdentifierReference }
	IfStatement                = func(t js.NodeType) bool { return t == js.IfStatement }
	ImportDeclaration          = func(t js.NodeType) bool { return t == js.ImportDeclaration }
	ImportSpecifier            = func(t js.NodeType) bool { return t == js.ImportSpecifier }
	IndexAccess                = func(t js.NodeType) bool { return t == js.IndexAccess }
	IndexSignature             = func(t js.NodeType) bool { return t == js.IndexSignature }
	IndexedAccessType          = func(t js.NodeType) bool { return t == js.IndexedAccessType }
	Initializer                = func(t js.NodeType) bool { return t == js.Initializer }
	IntersectionType           = func(t js.NodeType) bool { return t == js.IntersectionType }
	JSXAttributeName           = func(t js.NodeType) bool { return t == js.JSXAttributeName }
	JSXClosingElement          = func(t js.NodeType) bool { return t == js.JSXClosingElement }
	JSXElement                 = func(t js.NodeType) bool { return t == js.JSXElement }
	JSXElementName             = func(t js.NodeType) bool { return t == js.JSXElementName }
	JSXExpression              = func(t js.NodeType) bool { return t == js.JSXExpression }
	JSXLiteral                 = func(t js.NodeType) bool { return t == js.JSXLiteral }
	JSXNormalAttribute         = func(t js.NodeType) bool { return t == js.JSXNormalAttribute }
	JSXOpeningElement          = func(t js.NodeType) bool { return t == js.JSXOpeningElement }
	JSXSelfClosingElement      = func(t js.NodeType) bool { return t == js.JSXSelfClosingElement }
	JSXSpreadAttribute         = func(t js.NodeType) bool { return t == js.JSXSpreadAttribute }
	JSXSpreadExpression        = func(t js.NodeType) bool { return t == js.JSXSpreadExpression }
	JSXText                    = func(t js.NodeType) bool { return t == js.JSXText }
	KeyOfType                  = func(t js.NodeType) bool { return t == js.KeyOfType }
	LabelIdentifier            = func(t js.NodeType) bool { return t == js.LabelIdentifier }
	LabelledStatement          = func(t js.NodeType) bool { return t == js.LabelledStatement }
	LexicalBinding             = func(t js.NodeType) bool { return t == js.LexicalBinding }
	LexicalDeclaration         = func(t js.NodeType) bool { return t == js.LexicalDeclaration }
	Literal                    = func(t js.NodeType) bool { return t == js.Literal }
	LiteralPropertyName        = func(t js.NodeType) bool { return t == js.LiteralPropertyName }
	LiteralType                = func(t js.NodeType) bool { return t == js.LiteralType }
	LogicalANDExpression       = func(t js.NodeType) bool { return t == js.LogicalANDExpression }
	LogicalORExpression        = func(t js.NodeType) bool { return t == js.LogicalORExpression }
	MappedType                 = func(t js.NodeType) bool { return t == js.MappedType }
	MemberMethod               = func(t js.NodeType) bool { return t == js.MemberMethod }
	MemberVar                  = func(t js.NodeType) bool { return t == js.MemberVar }
	Method                     = func(t js.NodeType) bool { return t == js.Method }
	MethodSignature            = func(t js.NodeType) bool { return t == js.MethodSignature }
	Module                     = func(t js.NodeType) bool { return t == js.Module }
	ModuleSpecifier            = func(t js.NodeType) bool { return t == js.ModuleSpecifier }
	MultiplicativeExpression   = func(t js.NodeType) bool { return t == js.MultiplicativeExpression }
	NameSpaceImport            = func(t js.NodeType) bool { return t == js.NameSpaceImport }
	NamedImports               = func(t js.NodeType) bool { return t == js.NamedImports }
	NewExpression              = func(t js.NodeType) bool { return t == js.NewExpression }
	NewTarget                  = func(t js.NodeType) bool { return t == js.NewTarget }
	ObjectLiteral              = func(t js.NodeType) bool { return t == js.ObjectLiteral }
	ObjectPattern              = func(t js.NodeType) bool { return t == js.ObjectPattern }
	ObjectType                 = func(t js.NodeType) bool { return t == js.ObjectType }
	Parameters                 = func(t js.NodeType) bool { return t == js.Parameters }
	Parenthesized              = func(t js.NodeType) bool { return t == js.Parenthesized }
	ParenthesizedType          = func(t js.NodeType) bool { return t == js.ParenthesizedType }
	PostDec                    = func(t js.NodeType) bool { return t == js.PostDec }
	PostInc                    = func(t js.NodeType) bool { return t == js.PostInc }
	PreDec                     = func(t js.NodeType) bool { return t == js.PreDec }
	PreInc                     = func(t js.NodeType) bool { return t == js.PreInc }
	PredefinedType             = func(t js.NodeType) bool { return t == js.PredefinedType }
	Property                   = func(t js.NodeType) bool { return t == js.Property }
	PropertyAccess             = func(t js.NodeType) bool { return t == js.PropertyAccess }
	PropertyBinding            = func(t js.NodeType) bool { return t == js.PropertyBinding }
	PropertySignature          = func(t js.NodeType) bool { return t == js.PropertySignature }
	Readonly                   = func(t js.NodeType) bool { return t == js.Readonly }
	Regexp                     = func(t js.NodeType) bool { return t == js.Regexp }
	RelationalExpression       = func(t js.NodeType) bool { return t == js.RelationalExpression }
	RestParameter              = func(t js.NodeType) bool { return t == js.RestParameter }
	ReturnStatement            = func(t js.NodeType) bool { return t == js.ReturnStatement }
	Setter                     = func(t js.NodeType) bool { return t == js.Setter }
	ShiftExpression            = func(t js.NodeType) bool { return t == js.ShiftExpression }
	ShorthandProperty          = func(t js.NodeType) bool { return t == js.ShorthandProperty }
	SingleNameBinding          = func(t js.NodeType) bool { return t == js.SingleNameBinding }
	SpreadElement              = func(t js.NodeType) bool { return t == js.SpreadElement }
	SpreadProperty             = func(t js.NodeType) bool { return t == js.SpreadProperty }
	Static                     = func(t js.NodeType) bool { return t == js.Static }
	SuperExpression            = func(t js.NodeType) bool { return t == js.SuperExpression }
	SwitchStatement            = func(t js.NodeType) bool { return t == js.SwitchStatement }
	SyntaxProblem              = func(t js.NodeType) bool { return t == js.SyntaxProblem }
	TaggedTemplate             = func(t js.NodeType) bool { return t == js.TaggedTemplate }
	TemplateLiteral            = func(t js.NodeType) bool { return t == js.TemplateLiteral }
	This                       = func(t js.NodeType) bool { return t == js.This }
	ThisType                   = func(t js.NodeType) bool { return t == js.ThisType }
	ThrowStatement             = func(t js.NodeType) bool { return t == js.ThrowStatement }
	TryStatement               = func(t js.NodeType) bool { return t == js.TryStatement }
	TsAmbientBinding           = func(t js.NodeType) bool { return t == js.TsAmbientBinding }
	TsAmbientClass             = func(t js.NodeType) bool { return t == js.TsAmbientClass }
	TsAmbientClassBody         = func(t js.NodeType) bool { return t == js.TsAmbientClassBody }
	TsAmbientEnum              = func(t js.NodeType) bool { return t == js.TsAmbientEnum }
	TsAmbientFunction          = func(t js.NodeType) bool { return t == js.TsAmbientFunction }
	TsAmbientFunctionMember    = func(t js.NodeType) bool { return t == js.TsAmbientFunctionMember }
	TsAmbientImportAlias       = func(t js.NodeType) bool { return t == js.TsAmbientImportAlias }
	TsAmbientIndexMember       = func(t js.NodeType) bool { return t == js.TsAmbientIndexMember }
	TsAmbientInterface         = func(t js.NodeType) bool { return t == js.TsAmbientInterface }
	TsAmbientModule            = func(t js.NodeType) bool { return t == js.TsAmbientModule }
	TsAmbientNamespace         = func(t js.NodeType) bool { return t == js.TsAmbientNamespace }
	TsAmbientPropertyMember    = func(t js.NodeType) bool { return t == js.TsAmbientPropertyMember }
	TsAmbientTypeAlias         = func(t js.NodeType) bool { return t == js.TsAmbientTypeAlias }
	TsAmbientVar               = func(t js.NodeType) bool { return t == js.TsAmbientVar }
	TsAsExpression             = func(t js.NodeType) bool { return t == js.TsAsExpression }
	TsCastExpression           = func(t js.NodeType) bool { return t == js.TsCastExpression }
	TsEnum                     = func(t js.NodeType) bool { return t == js.TsEnum }
	TsEnumBody                 = func(t js.NodeType) bool { return t == js.TsEnumBody }
	TsEnumMember               = func(t js.NodeType) bool { return t == js.TsEnumMember }
	TsExportAssignment         = func(t js.NodeType) bool { return t == js.TsExportAssignment }
	TsImplementsClause         = func(t js.NodeType) bool { return t == js.TsImplementsClause }
	TsImportAliasDeclaration   = func(t js.NodeType) bool { return t == js.TsImportAliasDeclaration }
	TsImportRequireDeclaration = func(t js.NodeType) bool { return t == js.TsImportRequireDeclaration }
	TsIndexMemberDeclaration   = func(t js.NodeType) bool { return t == js.TsIndexMemberDeclaration }
	TsInterface                = func(t js.NodeType) bool { return t == js.TsInterface }
	TsInterfaceExtends         = func(t js.NodeType) bool { return t == js.TsInterfaceExtends }
	TsNamespace                = func(t js.NodeType) bool { return t == js.TsNamespace }
	TsNamespaceBody            = func(t js.NodeType) bool { return t == js.TsNamespaceBody }
	TsNonNull                  = func(t js.NodeType) bool { return t == js.TsNonNull }
	TsThisParameter            = func(t js.NodeType) bool { return t == js.TsThisParameter }
	TupleType                  = func(t js.NodeType) bool { return t == js.TupleType }
	TypeAliasDeclaration       = func(t js.NodeType) bool { return t == js.TypeAliasDeclaration }
	TypeAnnotation             = func(t js.NodeType) bool { return t == js.TypeAnnotation }
	TypeArguments              = func(t js.NodeType) bool { return t == js.TypeArguments }
	TypeConstraint             = func(t js.NodeType) bool { return t == js.TypeConstraint }
	TypeName                   = func(t js.NodeType) bool { return t == js.TypeName }
	TypeParameter              = func(t js.NodeType) bool { return t == js.TypeParameter }
	TypeParameters             = func(t js.NodeType) bool { return t == js.TypeParameters }
	TypePredicate              = func(t js.NodeType) bool { return t == js.TypePredicate }
	TypeQuery                  = func(t js.NodeType) bool { return t == js.TypeQuery }
	TypeReference              = func(t js.NodeType) bool { return t == js.TypeReference }
	UnaryExpression            = func(t js.NodeType) bool { return t == js.UnaryExpression }
	UnionType                  = func(t js.NodeType) bool { return t == js.UnionType }
	VariableDeclaration        = func(t js.NodeType) bool { return t == js.VariableDeclaration }
	VariableStatement          = func(t js.NodeType) bool { return t == js.VariableStatement }
	WhileStatement             = func(t js.NodeType) bool { return t == js.WhileStatement }
	WithStatement              = func(t js.NodeType) bool { return t == js.WithStatement }
	Yield                      = func(t js.NodeType) bool { return t == js.Yield }
	NoSubstitutionTemplate     = func(t js.NodeType) bool { return t == js.NoSubstitutionTemplate }
	TemplateHead               = func(t js.NodeType) bool { return t == js.TemplateHead }
	TemplateMiddle             = func(t js.NodeType) bool { return t == js.TemplateMiddle }
	TemplateTail               = func(t js.NodeType) bool { return t == js.TemplateTail }
	BindingPattern             = OneOf(js.BindingPattern...)
	CaseClause                 = OneOf(js.CaseClause...)
	ClassElement               = OneOf(js.ClassElement...)
	Declaration                = OneOf(js.Declaration...)
	ElementPattern             = OneOf(js.ElementPattern...)
	ExportElement              = OneOf(js.ExportElement...)
	Expression                 = OneOf(js.Expression...)
	JSXAttribute               = OneOf(js.JSXAttribute...)
	JSXAttributeValue          = OneOf(js.JSXAttributeValue...)
	JSXChild                   = OneOf(js.JSXChild...)
	MethodDefinition           = OneOf(js.MethodDefinition...)
	Modifier                   = OneOf(js.Modifier...)
	ModuleItem                 = OneOf(js.ModuleItem...)
	NamedImport                = OneOf(js.NamedImport...)
	Parameter                  = OneOf(js.Parameter...)
	PropertyDefinition         = OneOf(js.PropertyDefinition...)
	PropertyName               = OneOf(js.PropertyName...)
	PropertyPattern            = OneOf(js.PropertyPattern...)
	Statement                  = OneOf(js.Statement...)
	StatementListItem          = OneOf(js.StatementListItem...)
	TokenSet                   = OneOf(js.TokenSet...)
	TsAmbientClassElement      = OneOf(js.TsAmbientClassElement...)
	TsAmbientElement           = OneOf(js.TsAmbientElement...)
	TsType                     = OneOf(js.TsType...)
	TypeMember                 = OneOf(js.TypeMember...)
)

func OneOf(types ...js.NodeType) Selector {
	if len(types) == 0 {
		return func(js.NodeType) bool { return false }
	}
	const bits = 32
	max := 1
	for _, t := range types {
		if int(t) > max {
			max = int(t)
		}
	}
	size := (max + bits) / bits
	bitarr := make([]uint32, size)
	for _, t := range types {
		bitarr[uint(t)/bits] |= 1 << (uint(t) % bits)
	}
	return func(t js.NodeType) bool {
		i := uint(t) / bits
		return int(i) < len(bitarr) && bitarr[i]&(1<<(uint(t)%bits)) != 0
	}
}
