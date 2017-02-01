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
	Type string
	Statements []Node
}

func (bs *BlockStatement) statementNode() {

}

func (bs *BlockStatement) String() string {

	var out bytes.Buffer

	//out.WriteString("["+bs.Type+"]: ")

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type ExpressionStatement struct {
	Type string
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {

}

func (es *ExpressionStatement) String() string {

	return es.Expression.String()+";";
}


/** Expressions **/

type Prefix struct {
	Type 	 string
	Operator string
	Right    Expression
}

func (pe *Prefix) expressionNode() {}

func (pe *Prefix) String() string {

	var out bytes.Buffer

	//out.WriteString("["+pe.Type+"]: ")

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type Infix struct {
	Type 	 string
	Left     Expression
	Operator string
	Right    Expression
}

func (oe *Infix) expressionNode() {}

func (oe *Infix) String() string {

	var out bytes.Buffer

	//out.WriteString("["+oe.Type+"]: ")

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}

type Identifier struct {
	Type string
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) String() string {

	//return "["+i.Type+"]: "+i.Value
	return i.Value
}


type Boolean struct {
	Type string
	Value bool
}

func (b *Boolean) expressionNode() {}

func (b *Boolean) String() string {

	//return "["+b.Type+"]: "+strconv.FormatBool(b.Value);
	return strconv.FormatBool(b.Value);
}


type IntegerLiteral struct {
	Type string
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) String() string {

	//return "["+il.Type+"]: "+strconv.FormatInt(il.Value, 10);
	return strconv.FormatInt(il.Value, 10);
}


type FloatLiteral struct {
	Type string
	Value float64
}

func (f *FloatLiteral) expressionNode() {}

func (f *FloatLiteral) String() string {

	return strconv.FormatFloat(f.Value, 'E', -1, 64)
}

type String struct {
	Type string
	Value string
}


func (s *String) expressionNode() {}

func (s *String) String() string {

	return s.Value
}
