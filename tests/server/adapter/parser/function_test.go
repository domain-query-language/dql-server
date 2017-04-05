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
			[]ast.Parameter{},
			ast.BlockStatement{
				ast.BLOCK_STATEMENT,
				[]ast.Node,
			},
		},
	},
	/*{
		`function doThing() {
			return a;
		}`,
	},
	{
		`function doThing(value\name n, value\age a) {
			return a;
		}`,
	},
	*/
}

func TestFunctionParsing(t *testing.T) {

	for _, testCase := range functions {

		p := parser.NewFunction(testCase);

		parsed, err := p.ParseFunction();

		if (err != nil) {
			t.Error("Unexpected error while parsing: "+testCase.function)
			t.Error("Err: "+err)
		}

		if (!parsed.String != testCase.node.String()) {
			t.Error("Function parsed incorrectly: "+testCase.function)
			t.Error("Expected: "+testCase.node.String())
			t.Error("Actual: "+parsed.String())
		}
	}
}

/*
Todo
1. check
2. when
3. properties
4. handler
 */

