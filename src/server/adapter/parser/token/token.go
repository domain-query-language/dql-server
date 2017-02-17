package token

import (
	"fmt"
)

type Token struct {
	Type TokenType
	Val  string
	Pos  int
}

func NewToken(typ TokenType, val string, pos int) Token {
	return Token{typ, val, pos};
}

const IgnoreTokenPos = -1

func (t *Token) Compare(o Token) bool{
	if (t.Pos == IgnoreTokenPos || o.Pos == IgnoreTokenPos) {
		return t.Type == o.Type && t.Val == o.Val;
	}
	return t.Type == o.Type && t.Val == o.Val && t.Pos == o.Pos;
}

func (i *Token) String() string {
	switch i.Type {
	case EOF:
		return "EOF"
	}
	val := i.Val

	return fmt.Sprintf("Token(%v, %q, %d)", i.Type, val, i.Pos)
}

type TokenType string

const (
	EOF TokenType	= "eof"
	ERROR 		= "error"

	CLASSOPEN 	= "<|"
	CLASSCLOSE 	= "|>"
	OBJECTNAME 	= "objectName"

	//DQL Keywords - Objects
	CREATE 	   = "create"
	LIST	   = "list"
	AS 	   = "as"
	ON 	   = "on"
	USING	   = "using"
	FOR   	   = "for"
	IN	   = "in"
	WITHIN	   = "with"

	// Class components
	PROPERTIES = "properties"
	CHECK      = "check"
	HANDLER    = "handler"
	FUNCTION   = "function"
	WHEN  = "when"

	// Command Handler statements
	ASSERT  = "assert"
	NOT 	= "not"
	RUN 	= "run"
	APPLY 	= "apply"

	// Operators
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"
	REMAINDER = "%"
	ARROW 	  = "->"
	STRONGARROW = "=>"
	AND 	  = "and"
	OR 	  = "or"
	LT 	  = "<"
	GT 	  = ">"
	EQ 	  = "=="
	NOTEQ	  = "!="
	LTOREQ 	  = "<="
	GTOREQ 	  = ">="

	// Delimiters
	COMMA    = ","
	SEMICOLON= ";"
	COLON    = ":"
	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	//Types
	INTEGER	= "integer"
	FLOAT   = "float"
	BOOLEAN = "boolean"
	STRING  = "string"
	NULL    = "null"
	IDENT 	= "identifier"

	//Statements
	IF 	= "if"
	ELSEIF 	= "else if"
	ELSE 	= "else"
	RETURN 	= "return"
	FOREACH = "foreach"
)

func Semicolon(pos int) Token {
	return NewToken(SEMICOLON, ";", pos);
}

func ClsOpen(pos int) Token {
	return NewToken(CLASSOPEN, "<|", pos);
}

func ClsClose(pos int) Token {
	return NewToken(CLASSCLOSE, "|>", pos);
}