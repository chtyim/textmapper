// generated by Textmapper; DO NOT EDIT

package ast

import (
	"github.com/inspirer/textmapper/tm-go/parsers/tm"
	"log"
)

func ToTmNode(n Node) TmNode {
	if n == nil {
		return nil
	}
	switch n.Type() {
	case tm.KeyValue:
		return &KeyValue{n}
	case tm.AnnotationImpl:
		return &AnnotationImpl{n}
	case tm.Annotations:
		return &Annotations{n}
	case tm.ArgumentFalse:
		return &ArgumentFalse{n}
	case tm.ArgumentImpl:
		return &ArgumentImpl{n}
	case tm.ArgumentTrue:
		return &ArgumentTrue{n}
	case tm.Array:
		return &Array{n}
	case tm.Assoc:
		return &Assoc{n}
	case tm.BooleanLiteral:
		return &BooleanLiteral{n}
	case tm.Command:
		return &Command{n}
	case tm.DirectiveAssert:
		return &DirectiveAssert{n}
	case tm.DirectiveBrackets:
		return &DirectiveBrackets{n}
	case tm.DirectiveInput:
		return &DirectiveInput{n}
	case tm.DirectivePrio:
		return &DirectivePrio{n}
	case tm.DirectiveSet:
		return &DirectiveSet{n}
	case tm.ExclusiveStates:
		return &ExclusiveStates{n}
	case tm.GrammarParts:
		return &GrammarParts{n}
	case tm.Header:
		return &Header{n}
	case tm.Identifier:
		return &Identifier{n}
	case tm.Import:
		return &Import{n}
	case tm.InclusiveStates:
		return &InclusiveStates{n}
	case tm.InlineParameter:
		return &InlineParameter{n}
	case tm.Input:
		return &Input{n}
	case tm.Inputref:
		return &Inputref{n}
	case tm.IntegerLiteral:
		return &IntegerLiteral{n}
	case tm.InterfaceType:
		return &InterfaceType{n}
	case tm.Lexeme:
		return &Lexeme{n}
	case tm.LexemeAttribute:
		return &LexemeAttribute{n}
	case tm.LexemeAttrs:
		return &LexemeAttrs{n}
	case tm.LexerState:
		return &LexerState{n}
	case tm.ListSeparator:
		return &ListSeparator{n}
	case tm.Name:
		return &Name{n}
	case tm.NamedPattern:
		return &NamedPattern{n}
	case tm.Nonterm:
		return &Nonterm{n}
	case tm.NontermParams:
		return &NontermParams{n}
	case tm.ParamModifier:
		return &ParamModifier{n}
	case tm.ParamRef:
		return &ParamRef{n}
	case tm.ParamType:
		return &ParamType{n}
	case tm.Pattern:
		return &Pattern{n}
	case tm.Predicate:
		return &Predicate{n}
	case tm.PredicateAnd:
		return &PredicateAnd{n}
	case tm.PredicateEq:
		return &PredicateEq{n}
	case tm.PredicateNot:
		return &PredicateNot{n}
	case tm.PredicateNotEq:
		return &PredicateNotEq{n}
	case tm.PredicateOr:
		return &PredicateOr{n}
	case tm.RawType:
		return &RawType{n}
	case tm.References:
		return &References{n}
	case tm.RhsAnnotated:
		return &RhsAnnotated{n}
	case tm.RhsAssignment:
		return &RhsAssignment{n}
	case tm.RhsCast:
		return &RhsCast{n}
	case tm.RhsIgnored:
		return &RhsIgnored{n}
	case tm.RhsNested:
		return &RhsNested{n}
	case tm.RhsOptional:
		return &RhsOptional{n}
	case tm.RhsPlusAssignment:
		return &RhsPlusAssignment{n}
	case tm.RhsPlusList:
		return &RhsPlusList{n}
	case tm.RhsPrimary:
		return &RhsPrimary{n}
	case tm.RhsQuantifier:
		return &RhsQuantifier{n}
	case tm.RhsSet:
		return &RhsSet{n}
	case tm.RhsStarList:
		return &RhsStarList{n}
	case tm.RhsSuffix:
		return &RhsSuffix{n}
	case tm.RhsSymbol:
		return &RhsSymbol{n}
	case tm.Rule:
		return &Rule{n}
	case tm.RuleAction:
		return &RuleAction{n}
	case tm.SetAnd:
		return &SetAnd{n}
	case tm.SetComplement:
		return &SetComplement{n}
	case tm.SetCompound:
		return &SetCompound{n}
	case tm.SetOr:
		return &SetOr{n}
	case tm.SetSymbol:
		return &SetSymbol{n}
	case tm.StateSelector:
		return &StateSelector{n}
	case tm.Stateref:
		return &Stateref{n}
	case tm.StringLiteral:
		return &StringLiteral{n}
	case tm.SubType:
		return &SubType{n}
	case tm.Symref:
		return &Symref{n}
	case tm.SymrefArgs:
		return &SymrefArgs{n}
	case tm.SyntaxProblem:
		return &SyntaxProblem{n}
	case tm.TemplateParam:
		return &TemplateParam{n}
	case tm.VoidType:
		return &VoidType{n}
	case tm.InvalidToken, tm.MultilineComment, tm.Comment:
		return &Token{n}
	}
	log.Fatalf("unknown node type %v\n", n.Type())
	return nil
}
