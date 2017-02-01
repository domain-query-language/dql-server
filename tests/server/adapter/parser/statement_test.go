package parser

import (
	"testing"
	"github.com/domain-query-language/dql-server/src/server/adapter/ast"
	"github.com/domain-query-language/dql-server/src/server/adapter/parser"
)


func stmtBlk(node ast.Node) *ast.BlockStatement {

	return &ast.BlockStatement{
		Type: "blockstatement",
		Statements: []ast.Node{node},
	}
}

func compareBlockStatements(a *ast.BlockStatement, b *ast.BlockStatement) bool {

	return a.String() == b.String();
}



var prefixExpressions = []struct {
	expression string
	node ast.Node
}{
	{
		"--a;",
		&ast.Prefix{
			"prefix",
			"--",
			&ast.Identifier{
				"identifier",
				"a",
			},
		},
	},{
		"++a;",
		&ast.Prefix{
			"prefix",
			"++",
			&ast.Identifier{
				"identifier",
				"a",
			},
		},
	},{
		"!true;",
		&ast.Prefix{
			"prefix",
			"!",
			&ast.Identifier{
				"boolean",
				"true",
			},
		},
	},{
		"-15;",
		&ast.Prefix{
			"prefix",
			"-",
			&ast.Identifier{
				"integer",
				"15",
			},
		},
	},
}

func TestPrefixExpressions(t *testing.T) {

	for _, testCase := range prefixExpressions {

		p := parser.NewStatement(testCase.expression);

		actual, err := p.ParseBlockStatement();
		expected := stmtBlk(testCase.node);

		if (err != nil) {
			t.Error("Got error on '"+testCase.expression+"'");
			t.Error(err.Error());
		}

		if (!compareBlockStatements(actual, expected)) {
			t.Error("Expected AST does not match actual");
			t.Error("Expected: "+expected.String());
			t.Error("Actual: "+actual.String());
		}
	}
}

var infixExpressions = []struct {
	expression string
	node ast.Node
}{
	{"5 + 5;", infixInt(5, "+", 5)},
	{"5 - 5;", infixInt(5, "-", 5)},
	{"5 * 5;", infixInt(5, "*", 5)},
	{"5 / 5;", infixInt(5, "/", 5)},
	{"5 > 5;", infixInt(5, ">", 5)},
	{"5 < 5;", infixInt(5, "<", 5)},
	{"5 == 5;", infixInt(5, "==", 5)},
	{"5 != 5;", infixInt(5, "!=", 5)},
	{"foobar + barfoo;", infixIdent("foobar", "+", "barfoo")},
	{"foobar - barfoo;", infixIdent("foobar", "-", "barfoo")},
	{"foobar * barfoo;", infixIdent("foobar", "*", "barfoo")},
	{"foobar / barfoo;", infixIdent("foobar", "/", "barfoo")},
	{"foobar > barfoo;", infixIdent("foobar", ">", "barfoo")},
	{"foobar < barfoo;", infixIdent("foobar", "<", "barfoo")},
	{"foobar == barfoo;", infixIdent("foobar", "==", "barfoo")},
	{"foobar != barfoo;", infixIdent("foobar", "!=", "barfoo")},
	{"true == true;", infixBool(true, "==", true)},
	{"true != false;", infixBool(true, "!=", false)},
	{"false == false;", infixBool(false, "==", false)},
}

func TestInfixExpressions(t *testing.T) {

	for _, testCase := range infixExpressions {

		p := parser.NewStatement(testCase.expression);

		actual, err := p.ParseBlockStatement();
		expected := stmtBlk(testCase.node);

		if (err != nil) {
			t.Error("Got error on '"+testCase.expression+"'");
			t.Error(err.Error());
		}

		if (!compareBlockStatements(actual, expected)) {
			t.Error("Expected AST does not match actual");
			t.Error("Expected: "+expected.String());
			t.Error("Actual: "+actual.String());
		}
	}
}

func infixInt(a int64, op string, b int64) ast.Node {

	return &ast.Infix{
		"infix",
		&ast.IntegerLiteral{"integer",a},
		op,
		&ast.IntegerLiteral{"integer",b},
	};
}

func infixIdent(a string, op string, b string) ast.Node {

	return &ast.Infix{
		"infix",
		&ast.Identifier{"identifier",a},
		op,
		&ast.Identifier{"identifier",b},
	};
}

func infixBool(a bool, op string, b bool) ast.Node {

	return &ast.Infix{
		"infix",
		&ast.Boolean{"boolean",a},
		op,
		&ast.Boolean{"boolean",b},
	};
}