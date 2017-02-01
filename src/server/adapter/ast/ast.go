package ast

import (
	"bytes"
	"strconv"
)

// The base Node interface
type Node interface {
	String() string
}

// All statement nodes implement this
type Statement interface {
	Node
	statementNode()
}

// All expression nodes implement this
type Expression interface {
	Node
	expressionNode()
}


/** Statements **/

type BlockStatement struct {
	Statements []Node
}

func (bs *BlockStatement) statementNode() {

}

func (bs *BlockStatement) String() string {

	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}


/** Expressions **/

type Prefix struct {
	Type string
	Operator string
	Right    Expression
}

func (pe *Prefix) expressionNode() {}

func (pe *Prefix) String() string {

	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type Identifier struct {
	Type string
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) String() string {

	return i.Value
}


type Boolean struct {
	Type string
	Value bool
}

func (b *Boolean) expressionNode() {}

func (b *Boolean) String() string {

	return strconv.FormatBool(b.Value);
}


type IntegerLiteral struct {
	Type string
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) String() string {

	return strconv.FormatInt(il.Value, 10);
}