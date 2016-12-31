// generated by Textmapper; DO NOT EDIT

package tm

import (
	"fmt"
)

// Token is an enum of all terminal symbols of the tm language.
type Token int

// Token values.
const (
	UNAVAILABLE Token = iota-1
	EOI

	REGEXP
	SCON
	ICON
	_SKIP
	_SKIP_COMMENT
	_SKIP_MULTILINE
	REM // %
	COLONCOLONASSIGN // ::=
	COLONCOLON // ::
	OR // |
	OROR // ||
	ASSIGN // =
	ASSIGNASSIGN // ==
	EXCLASSIGN // !=
	ASSIGNGT // =>
	SEMICOLON // ;
	DOT // .
	COMMA // ,
	COLON // :
	LBRACK // [
	RBRACK // ]
	LPAREN // (
	RPAREN // )
	LBRACETILDE // {~
	RBRACE // }
	LT // <
	GT // >
	MULT // *
	PLUS // +
	PLUSASSIGN // +=
	QUEST // ?
	EXCL // !
	TILDE // ~
	AND // &
	ANDAND // &&
	DOLLAR // $
	ATSIGN // @
	ERROR
	ID
	TRUE // true
	FALSE // false
	NEW // new
	SEPARATOR // separator
	AS // as
	IMPORT // import
	SET // set
	BRACKETS // brackets
	INLINE // inline
	PREC // prec
	SHIFT // shift
	RETURNS // returns
	INPUT // input
	LEFT // left
	RIGHT // right
	NONASSOC // nonassoc
	GENERATE // generate
	ASSERT // assert
	EMPTY // empty
	NONEMPTY // nonempty
	GLOBAL // global
	EXPLICIT // explicit
	LOOKAHEAD // lookahead
	PARAM // param
	FLAG // flag
	NOMINUSEOI // no-eoi
	SOFT // soft
	CLASS // class
	INTERFACE // interface
	VOID // void
	SPACE // space
	LAYOUT // layout
	LANGUAGE // language
	LALR // lalr
	LEXER // lexer
	PARSER // parser
	REDUCE // reduce
	CODE // {
	LBRACE // {

	NumTokens
)

var tokenStr = [...]string{
	"EOI",

	"REGEXP",
	"SCON",
	"ICON",
	"_SKIP",
	"_SKIP_COMMENT",
	"_SKIP_MULTILINE",
	"%",
	"::=",
	"::",
	"|",
	"||",
	"=",
	"==",
	"!=",
	"=>",
	";",
	".",
	",",
	":",
	"[",
	"]",
	"(",
	")",
	"{~",
	"}",
	"<",
	">",
	"*",
	"+",
	"+=",
	"?",
	"!",
	"~",
	"&",
	"&&",
	"$",
	"@",
	"ERROR",
	"ID",
	"true",
	"false",
	"new",
	"separator",
	"as",
	"import",
	"set",
	"brackets",
	"inline",
	"prec",
	"shift",
	"returns",
	"input",
	"left",
	"right",
	"nonassoc",
	"generate",
	"assert",
	"empty",
	"nonempty",
	"global",
	"explicit",
	"lookahead",
	"param",
	"flag",
	"no-eoi",
	"soft",
	"class",
	"interface",
	"void",
	"space",
	"layout",
	"language",
	"lalr",
	"lexer",
	"parser",
	"reduce",
	"{",
	"{",
}

func (tok Token) String() string {
	if tok >= 0 && int(tok) < len(tokenStr) {
		return tokenStr[tok]
	}
	return fmt.Sprintf("token(%d)", tok)
}
