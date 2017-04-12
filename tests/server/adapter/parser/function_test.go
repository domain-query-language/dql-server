package parser

import (
	"github.com/domain-query-language/dql-server/src/server/adapter/ast"
	"testing"
	"github.com/domain-query-language/dql-server/src/server/adapter/parser"
)

type funcTestCase struct {
	function string
	node ast.Node
}

func testParsingFunction(testCase funcTestCase, t *testing.T) {

	p := parser.NewFunction(testCase.function);

	parsed, err := p.ParseObjectComponent();

	if (err != nil) {
		t.Error("Unexpected error while parsing: "+testCase.function)
		t.Error("Err: "+err.Error())
	}

	if (parsed.String() != testCase.node.String()) {
		t.Error("Function parsed incorrectly: "+testCase.function)
		t.Error("Expected: "+testCase.node.String())
		t.Error("Actual: "+parsed.String())
	}
}

var functions = []funcTestCase {
	{
		`function doThing() {

		}`,
		&ast.Function{
			ast.FUNCTION,
			"doThing",
			[]*ast.Parameter{},
			&ast.BlockStatement{
				ast.BLOCK_STATEMENT,
				[]ast.Node{},
			},
		},
	},
	{
		`function doThing2(value\name n, value\age a) {
			a;
		}`,
		&ast.Function{
			ast.FUNCTION,
			"doThing2",
			[]*ast.Parameter{
				{`value\name`, "n"},
				{`value\age`, "a"},
			},
			&ast.BlockStatement{
				ast.BLOCK_STATEMENT,
				[]ast.Node{
					&ast.ExpressionStatement{
						ast.EXPRESSION_STATEMENT,
						&ast.Identifier{
							ast.IDENTIFIER,
							"a",
						},
					},
				},
			},
		},
	},
}

func TestFunctionParsing(t *testing.T) {

	for _, testCase := range functions {

		testParsingFunction(testCase, t)
	}
}

func TestParsingMultipleFunctionsInOneString(t *testing.T) {

	multipleFunctions := ""

	for _, testCase := range functions {

		multipleFunctions += testCase.function
	}

	p := parser.NewFunction(multipleFunctions);

	for _, testCase := range functions {

		parsed, err := p.ParseObjectComponent();

		if (err != nil) {
			t.Error("Unexpected error while parsing: "+testCase.function)
			t.Error("Err: "+err.Error())
		}

		if (parsed.String() != testCase.node.String()) {
			t.Error("Function parsed incorrectly: "+testCase.function)
			t.Error("Expected: "+testCase.node.String())
			t.Error("Actual: "+parsed.String())
		}
	}
}

func TestPrintingFunction(t *testing.T) {

	input := `function a(value\b b,value\c c){a;return d; }`

	output := `function a(value\b b, value\c c) {
	a;
	return d;
}`;

	p := parser.NewFunction(input);

	parsed, err := p.ParseObjectComponent();

	if (err != nil) {
		t.Error("Unexpected error while parsing: "+input)
		t.Error("Err: "+err.Error())
	}

	if (parsed.String() != output) {
		t.Error("Function parsed incorrectly: "+input)
		t.Error("Expected: "+output)
		t.Error("Actual: "+parsed.String())
	}

}

var checkStatements = []funcTestCase{
	{
		`check (
			return a;
		)`,
		&ast.Function{
			ast.FUNCTION,
			"check",
			[]*ast.Parameter{},
			&ast.BlockStatement{
				ast.BLOCK_STATEMENT,
				[]ast.Node{
					&ast.Return{
						ast.RETURN_STATEMENT,
						&ast.Identifier{
							ast.IDENTIFIER,
							"a",
						},
					},
				},
			},
		},
	},
}

func TestCheckStatement(t *testing.T) {

	for _, testCase := range checkStatements {

		testParsingFunction(testCase, t)
	}
}

var handlerStatements = []funcTestCase{
	{
		`handler {
			return a;
		}`,
		&ast.Function{
			ast.FUNCTION,
			"handler",
			[]*ast.Parameter{},
			&ast.BlockStatement{
				ast.BLOCK_STATEMENT,
				[]ast.Node{
					&ast.Return{
						ast.RETURN_STATEMENT,
						&ast.Identifier{
							ast.IDENTIFIER,
							"a",
						},
					},
				},
			},
		},
	},
}

func TestHandlerStatement(t *testing.T) {

	for _, testCase := range handlerStatements {

		testParsingFunction(testCase, t)
	}
}

var whenStatements = []funcTestCase{
	{
		`WHEN event 'started' {
			return a;
		}`,
		&ast.When{
			ast.WHEN,
			"started",
			&ast.BlockStatement{
				ast.BLOCK_STATEMENT,
				[]ast.Node{
					&ast.Return{
						ast.RETURN_STATEMENT,
						&ast.Identifier{
							ast.IDENTIFIER,
							"a",
						},
					},
				},
			},
		},
	},
}

func TestWhenStatement(t *testing.T) {

	for _, testCase := range whenStatements {

		testParsingFunction(testCase, t)
	}
}


/*
Done
- parsing two functions in sequence
- function with params
- function with body
- printing a function
- check
- handler
- when

Inprogress
- properties

Todo
- Clean up handing over of state back to function parser
 */

