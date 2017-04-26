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

func testParsingObjectComponent(testCase funcTestCase, t *testing.T) {

	p := parser.NewObjectComponent(testCase.function);

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

func TestFunctions(t *testing.T) {

	for _, testCase := range functions {

		testParsingObjectComponent(testCase, t)
	}
}

func TestParsingMultipleFunctionsInOneString(t *testing.T) {

	multipleFunctions := ""

	for _, testCase := range functions {

		multipleFunctions += testCase.function
	}

	p := parser.NewObjectComponent(multipleFunctions);

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

	p := parser.NewObjectComponent(input);

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

var checks = []funcTestCase{
	{
		`check (
			return a;
		)`,
		&ast.Check{
			ast.CHECK,
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

func TestChecks(t *testing.T) {

	for _, testCase := range checks {

		testParsingObjectComponent(testCase, t)
	}
}

var handlers = []funcTestCase{
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

func TestHandlers(t *testing.T) {

	for _, testCase := range handlers {

		testParsingObjectComponent(testCase, t)
	}
}

var whens = []funcTestCase{
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

func TestWhens(t *testing.T) {

	for _, testCase := range whens {

		testParsingObjectComponent(testCase, t)
	}
}

var properties = []funcTestCase{
	{
		`properties{
			value\service_charge service_charge;
			value\category category;
		}`,
		&ast.Properties{
			ast.PROPERTIES,
			[]*ast.Property{
				{
					ast.PROPERTY,
					"value\\service_charge",
					"service_charge",
					nil,
				},
				{
					ast.PROPERTY,
					"value\\category",
					"category",
					nil,
				},
			},
		},
	},
	{
		`properties{
			value\category category = 'value\category'("category");
		}`,
		&ast.Properties{
			ast.PROPERTIES,
			[]*ast.Property{
				{
					ast.PROPERTY,
					"value\\category",
					"category",
					&ast.ExpressionStatement{
						Type: ast.EXPRESSION_STATEMENT,
						Expression: &ast.ObjectCreation{
							ast.OBJECT_CREATION,
							"value\\category",
							[]ast.Expression{
								&ast.String{ast.STRING, "category"},
							},
						},
					},
				},
			},
		},
	},
}

func TestProperties(t *testing.T) {

	for _, testCase := range properties {

		testParsingObjectComponent(testCase, t)
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
- properties
- Clean up handing over of state back to function parser

Todo
- Create AST for check
- ObjectComponent parser should be instantiable from tokenstream
 */

