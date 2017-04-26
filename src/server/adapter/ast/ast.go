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
	ASSERT_STATEMENT = "assertstatement"
	APPLY_STATEMENT = "applystatement"

	PREFIX = "prefix"
	INFIX = "infix"
	IDENTIFIER = "identifier"
	INTEGER = "integer"
	BOOLEAN = "boolean"
	FLOAT = "float"
	STRING = "string"
	METHOD_CALL = "methodcall"
	ARRAY = "array"
	ARRAY_ACCESS = "arrayaccess"
	OBJECT_CREATION = "objectcreation"
	RUN_QUERY = "runquery"

	FUNCTION = "function"
	CHECK = "check"
	WHEN = "when"
	PROPERTIES = "properties"
	PROPERTY = "property"
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

// All expression nodes implement this
type ObjectComponent interface {
	Node
	objectComponentNode()
}


/** Statements **/

type BlockStatement struct {
	Type string
	Statements []Node
}

func (bs *BlockStatement) statementNode() {}

func (bs *BlockStatement) String() string {

	var out bytes.Buffer

	for i, s := range bs.Statements {
		if (s != nil) {
			out.WriteString(s.String())
		}
		if (i+1!= len(bs.Statements)) {
			out.WriteString("\n")
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

	return "return "+r.Expression.String()+";"
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


/** Handler statements **/

type AssertStatement struct {
	Type string
	Operator string
	Event Expression
}

func (a *AssertStatement) statementNode() {}

func (a *AssertStatement) String() string {

	str := "assert invariant "

	if (a.Operator == "") {
		str += a.Operator+""
	}

	return str + a.Event.String()
}


type ApplyStatement struct {
	Type string
	Event Expression
}

func (a *ApplyStatement) statementNode() {}

func (a *ApplyStatement) String() string {

	str := "apply event "

	return str + a.Event.String()
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

type Array struct {
	Type string
	Elements []Expression
}

func (a *Array) expressionNode() {}

func (a *Array) String() string {

	elms := make([]string, len(a.Elements))
	for i, elm := range a.Elements {
		elms[i] = elm.String()
	}

	return "["+strings.Join(elms, ", ")+"]";
}

type ArrayAccess struct {
	Type string
	Left  Expression
	Offset Expression
}

func (m *ArrayAccess) expressionNode() {}

func (a *ArrayAccess) String() string {

	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(a.Left.String())
	out.WriteString("[")
	out.WriteString(a.Offset.String())
	out.WriteString("])")

	return out.String()
}

type ObjectCreation struct {
	Type string
	Name  string
	Arguments []Expression
}

func (o *ObjectCreation) expressionNode() {}

func (o *ObjectCreation) String() string {

	args := make([]string, len(o.Arguments))
	for i, arg := range o.Arguments {
		args[i] = arg.String()
	}

	return "'"+o.Name + "'("+strings.Join(args, ", ")+")";
}

type RunQuery struct {
	Type string
	Query Expression
}

func (r *RunQuery) expressionNode() {}

func (r *RunQuery) String() string {

	return "run query "+r.Query.String()
}

type Function struct {
	Type string
	Name string
	Parameters []*Parameter
	Body Statement
}

func (f *Function) objectComponentNode() {}

func (f *Function) String() string {

	params := make([]string, len(f.Parameters))
	for i, param := range f.Parameters {
		params[i] = param.String()
	}

	body := f.Body.String()

	body = "\t"+strings.Replace(body, "\n", "\n\t", -1)

	return "function "+f.Name+"("+strings.Join(params, ", ")+") {\n"+body+"\n}"
}

type Parameter struct {
	Type string
	Name string
}

func (p Parameter) String() string {
	return p.Type+" "+p.Name
}


type Check struct {
	Type string
	Body Statement
}

func (f *Check) objectComponentNode() {}

func (f *Check) String() string {

	body := f.Body.String()

	body = "\t"+strings.Replace(body, "\n", "\n\t", -1)

	return "check {\n"+body+"\n}"
}

type When struct {
	Type string
	Event string
	Body Statement
}

func (f *When) objectComponentNode() {}

func (f *When) String() string {

	body := f.Body.String()

	body = "\t"+strings.Replace(body, "\n", "\n\t", -1)

	return "when event  "+f.Event+ "{\n"+body+"\n}"
}

type Properties struct {
	Type string
	Properties []*Property
}

func (p *Properties) String() string {

	var props = make([]string, len(p.Properties))
	for _, prop := range p.Properties {
		props = append(props, prop.String())
	}

	body := "\t"+strings.Replace(strings.Join(props, "\n"), "\n", "\n\t", -1)

	return "properties {\n"+body+"}"
}

func (p *Properties) objectComponentNode() {}

type Property struct {
	Type string
	ValueType string
	Name string
	Exp Node
}

func (p *Property) String() string {

	left := p.ValueType+" "+p.Name;
	if (p.Exp == nil) {
		return left + ";";
	}

	return left+" = "+p.Exp.String();
}

func (p *Property) statementNode() {}

