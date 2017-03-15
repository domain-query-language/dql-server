package ast

import (
	"bytes"
	"strconv"
	"strings"
)

const (
	BLOCK_STATEMENT string = "blockstatement"
	EXPRESSION_STATEMENT = "expressionstatement"
	RETURN_STATEMENT = "returnstatement"
	FOREACH_STATEMENT = "foreachstatement"
	IF_STATEMENT = "ifstatement"

	PREFIX = "prefix"
	INFIX = "infix"
	IDENTIFIER = "identifier"
	INTEGER = "integer"
	BOOLEAN = "boolean"
	FLOAT = "float"
	STRING = "string"
	METHOD_CALL = "methodcall"
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

func (bs *BlockStatement) statementNode() {}

func (bs *BlockStatement) String() string {

	var out bytes.Buffer

	for _, s := range bs.Statements {
		if (s != nil) {
			out.WriteString(s.String())
		}
	}

	return out.String()
}

type ExpressionStatement struct {
	Type string
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) String() string {

	return es.Expression.String()+";";
}


type Return struct {
	Type string
	Expression Expression
}

func (r *Return) statementNode() {}

func (r *Return) String() string {

	return "return "+r.Expression.String()
}


type IfStatement struct {
	Type 	   string
	Test       Expression
	Consequent Statement
	Alternate  Statement
}

func (i *IfStatement) statementNode() {}

func (i *IfStatement) String() string {

	str := "if "+i.Test.String()+"{\n    "+i.Consequent.String()+"\n}";

	if (i.Alternate != nil) {
		str += " else {\n    "+i.Alternate.String()+"\n}";
	}
	return str;
}


type ForeachStatement struct {
	Type 	   string
	Collection Expression
	Key 	   *Identifier
	Value      *Identifier
	Body 	   Statement
}

func (f *ForeachStatement) statementNode() {}

func (f *ForeachStatement) String() string {

	as := "";
	if (f.Key != nil) {
		as = "as "+f.Key.String()+" "
	}

	str := "foreach ("+f.Collection.String()+" "+as+f.Value.String()+") {\n    "+f.Body.String()+"\n}";

	return str;
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


type Integer struct {
	Type string
	Value int64
}

func (il *Integer) expressionNode() {}

func (il *Integer) String() string {

	//return "["+il.Type+"]: "+strconv.FormatInt(il.Value, 10);
	return strconv.FormatInt(il.Value, 10);
}


type Float struct {
	Type string
	Value float64
}

func (f *Float) expressionNode() {}

func (f *Float) String() string {

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

type MethodCall struct {
	Type string
	Method  Expression
	Arguments []Expression
}

func (m *MethodCall) expressionNode() {}

func (m *MethodCall) String() string {

	args := make([]string, len(m.Arguments))
	for i, arg := range m.Arguments {
		args[i] = arg.String()
	}

	return m.Method.String() + "("+strings.Join(args, ", ")+")";
}
