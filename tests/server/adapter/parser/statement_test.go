package parser

import (
	"testing"
	"github.com/domain-query-language/dql-server/src/server/adapter/ast"
	"github.com/domain-query-language/dql-server/src/server/adapter/parser"
)

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

func stmtBlk(node ast.Node) *ast.BlockStatement {

	return &ast.BlockStatement{
		Statements: []ast.Node{node},
	}
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

func compareBlockStatements(a *ast.BlockStatement, b *ast.BlockStatement) bool {

	return a.String() == b.String();
}