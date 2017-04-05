package parser

import (
	"github.com/domain-query-language/dql-server/src/server/adapter/ast"
	"testing"
	"github.com/domain-query-language/dql-server/src/server/adapter/parser"
)

var functions = []struct{
	function string
	node ast.Node
}{
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

		p := parser.NewFunction(testCase.function);

		parsed, err := p.ParseFunction();

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

func TestParsingMultipleFunctionsInOneString(t *testing.T) {

	multipleFunctions := ""

	for _, testCase := range functions {

		multipleFunctions += testCase.function
	}

	p := parser.NewFunction(multipleFunctions);

	for _, testCase := range functions {

		parsed, err := p.ParseFunction();

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

	parsed, err := p.ParseFunction();

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

/*
Done
- parsing two functions in sequence
- function with params
- function with body
- printing a function

Todo
- check
- when
- properties
- handler
- Clean up handing over of state back to function parser
 */

