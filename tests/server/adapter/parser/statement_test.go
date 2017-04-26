package parser

import (
	"testing"
	"github.com/domain-query-language/dql-server/src/server/adapter/ast"
	"github.com/domain-query-language/dql-server/src/server/adapter/parser"
)


/*************************
  Helper funcs for tests
 ************************/

func compareBlockStatements(a *ast.BlockStatement, b *ast.BlockStatement) bool {

	return a.String() == b.String();
}


func blkStmt(nodes []ast.Node) *ast.BlockStatement {

	return &ast.BlockStatement{
		Type: ast.BLOCK_STATEMENT,
		Statements: nodes,
	}
}

func expStmt(exp ast.Expression) ast.Statement {

	return &ast.ExpressionStatement{
		Type: ast.EXPRESSION_STATEMENT,
		Expression: exp,
	}
}

func infixInt(a int64, op string, b int64) ast.Node {

	return expStmt(&ast.Infix{
		ast.INFIX,
		&ast.Integer{ast.INTEGER,a},
		op,
		&ast.Integer{ast.INTEGER,b},
	});
}

func infixIdent(a string, op string, b string) ast.Node {

	return expStmt(&ast.Infix{
		ast.INFIX,
		&ast.Identifier{ast.IDENTIFIER,a},
		op,
		&ast.Identifier{ast.IDENTIFIER,b},
	});
}

func infixBool(a bool, op string, b bool) ast.Node {

	return expStmt(&ast.Infix{
		ast.INFIX,
		&ast.Boolean{ast.BOOLEAN,a},
		op,
		&ast.Boolean{ast.BOOLEAN,b},
	});
}

type testCase struct {
	expression string
	node ast.Node
}

type testCases []testCase


func (testCases testCases) test(t *testing.T) {

	for _, testCase := range testCases {

		testCase.test(t)
	}
}

func (testCase testCase) test(t *testing.T) {

	statementBlockStr := "{ "+testCase.expression+" }";

	p := parser.NewStatement(statementBlockStr);

	actual, err := p.ParseBlockStatement();
	expected := blkStmt([]ast.Node{testCase.node});

	if (err != nil) {
		t.Error("Got error on '" + testCase.expression + "'");
		t.Error(err.Error());
	}

	if (actual!= nil && actual.Type != "blockstatement") {
		t.Error("Parser did not return blocsktatement");
	}

	if (!compareBlockStatements(actual, expected)) {
		t.Error("Expected AST does not match actual");
		t.Error("Expected: " + expected.String());
		t.Error("Actual: " + actual.String());
	}
}


/****************
  Test cases
 ***************/

var prefixExpressions = testCases{
	{
		"--a;",
		expStmt(&ast.Prefix{
			ast.PREFIX,
			"--",
			&ast.Identifier{
				ast.IDENTIFIER,
				"a",
			},
		}),
	}, {
		"++a;",
		expStmt(&ast.Prefix{
			ast.PREFIX,
			"++",
			&ast.Identifier{
				ast.IDENTIFIER,
				"a",
			},
		}),
	}, {
		"!true;",
		expStmt(&ast.Prefix{
			ast.PREFIX,
			"!",
			&ast.Identifier{
				ast.BOOLEAN,
				"true",
			},
		}),
	}, {
		"-15;",
		expStmt(&ast.Prefix{
			ast.PREFIX,
			"-",
			&ast.Identifier{
				ast.INTEGER,
				"15",
			},
		}),
	},
}

func TestPrefixExpressions(t *testing.T) {

	prefixExpressions.test(t)
}

var infixExpressions = testCases{
	{"5 + 5;", infixInt(5, "+", 5)},
	{"5 - 5;", infixInt(5, "-", 5)},
	{"5 * 5;", infixInt(5, "*", 5)},
	{"5 / 5;", infixInt(5, "/", 5)},
	{"5 > 5;", infixInt(5, ">", 5)},
	{"5 < 5;", infixInt(5, "<", 5)},
	{"5 == 5;", infixInt(5, "==", 5)},
	{"5 != 5;", infixInt(5, "!=", 5)},
	{"a + b;", infixIdent("a", "+", "b")},
	{"a - b;", infixIdent("a", "-", "b")},
	{"a * b;", infixIdent("a", "*", "b")},
	{"a / b;", infixIdent("a", "/", "b")},
	{"a > b;", infixIdent("a", ">", "b")},
	{"a < b;", infixIdent("a", "<", "b")},
	{"a == b;", infixIdent("a", "==", "b")},
	{"a != b;", infixIdent("a", "!=", "b")},
	{"a->b;", infixIdent("a", "->", "b")},
	{"a = b;", infixIdent("a", "=", "b")},
	{"true == true;", infixBool(true, "==", true)},
	{"true != false;", infixBool(true, "!=", false)},
	{"false == false;", infixBool(false, "==", false)},

}

func TestInfixExpressions(t *testing.T) {

	infixExpressions.test(t)
}

var basicTypes = testCases{
	{
		"15.1;",
		expStmt(&ast.Float{
			ast.FLOAT,
			15.1,
		}),
	},
	{
		"\"I am a string, ohh yes I am!\";",
		expStmt(&ast.String{
			ast.STRING,
			"I am a string, ohh yes I am!",
		}),
	},
}

func TestBasicTypes(t *testing.T) {

	basicTypes.test(t);
}

var methodCalls = testCases{
	{
		"a->b();",
		expStmt(&ast.MethodCall{
			ast.METHOD_CALL,
			&ast.Infix{
				ast.INFIX,
				&ast.Identifier{ast.IDENTIFIER, "a"},
				"->",
				&ast.Identifier{ast.IDENTIFIER, "b"},
			},
			[]ast.Expression{},
		}),
	},
	{
		"a->b(c, d);",
		expStmt(&ast.MethodCall{
			ast.METHOD_CALL,
			&ast.Infix{
				ast.INFIX,
				&ast.Identifier{ast.IDENTIFIER, "a"},
				"->",
				&ast.Identifier{ast.IDENTIFIER, "b"},
			},
			[]ast.Expression{
				&ast.Identifier{ast.IDENTIFIER, "c"},
				&ast.Identifier{ast.IDENTIFIER, "d"},
			},
		}),
	},
}

func TestMethodClass(t *testing.T) {

	methodCalls.test(t);
}

var arrayAccess = testCases{
	{
		"a[1];",
		expStmt(&ast.ArrayAccess{
			ast.ARRAY_ACCESS,
			&ast.Identifier{ast.IDENTIFIER, "a"},
			&ast.Integer{ast.INTEGER, 1},
		}),
	},
	{
		"a[b+c];",
		expStmt(&ast.ArrayAccess{
			ast.ARRAY_ACCESS,
			&ast.Identifier{ast.IDENTIFIER, "a"},
			&ast.Infix{
				ast.INFIX,
				&ast.Identifier{ast.IDENTIFIER, "b"},
				"+",
				&ast.Identifier{ast.IDENTIFIER,"c"},
			},
		}),
	},
}

func TestArrayAccess(t *testing.T) {

	arrayAccess.test(t);
}

var arrayCreation = testCases{
	{
		"[a];",
		expStmt(&ast.Array{
			ast.ARRAY,
			[]ast.Expression{
				&ast.Identifier{ast.IDENTIFIER, "a"},
			},
		}),
	},
	{
		"[a, b, 1];",
		expStmt(&ast.Array{
			ast.ARRAY,
			[]ast.Expression{
				&ast.Identifier{ast.IDENTIFIER, "a"},
				&ast.Identifier{ast.IDENTIFIER, "b"},
				&ast.Integer{ast.INTEGER, 1},
			},
		}),
	},
}

func TestArrayCreation(t *testing.T) {

	arrayCreation.test(t);
}

var statementBlock = testCase{
	`
	a;
	b;
	`,
	blkStmt([]ast.Node{
		expStmt(&ast.Identifier{
			ast.IDENTIFIER,
			"a",
		}),
		expStmt(&ast.Identifier{
			ast.IDENTIFIER,
			"b",
		}),
	}),
}

func TestBlockStatement(t *testing.T) {

	statementBlock.test(t);
}

var statements = testCases {
	{
		"return a;",
		blkStmt([]ast.Node{
			&ast.Return{
				ast.RETURN_STATEMENT,
				&ast.Identifier{
					ast.IDENTIFIER,
					"a",
				},
			},
		}),
	},
	{
		`if (a) { b; }`,
		blkStmt([]ast.Node{
			&ast.IfStatement{
				ast.IF_STATEMENT,
				&ast.Identifier{
					ast.IDENTIFIER,
					"a",
				},
				blkStmt([]ast.Node{
					expStmt(&ast.Identifier{
						ast.IDENTIFIER,
						"b",
					}),
				}),
				nil,
			},
		}),
	},
	{
		`if a {
			b;
		} else {
			c;
		}`,
		blkStmt([]ast.Node{
			&ast.IfStatement{
				ast.IF_STATEMENT,
				&ast.Identifier{
					ast.IDENTIFIER,
					"a",
				},
				blkStmt([]ast.Node{
					expStmt(&ast.Identifier{
						ast.IDENTIFIER,
						"b",
					}),
				}),
				blkStmt([]ast.Node{
					expStmt(&ast.Identifier{
						ast.IDENTIFIER,
						"c",
					}),
				}),
			},
		}),
	},
	{
		`if a {
			b;
		} else {
			if (c) {
				d;
			}
		}`,
		blkStmt([]ast.Node{
			&ast.IfStatement{
				ast.IF_STATEMENT,
				&ast.Identifier{
					ast.IDENTIFIER,
					"a",
				},
				blkStmt([]ast.Node{
					expStmt(&ast.Identifier{
						ast.IDENTIFIER,
						"b",
					}),
				}),
				blkStmt([]ast.Node{
					&ast.IfStatement{
						"if",
						&ast.Identifier{
							ast.IDENTIFIER,
							"c",
						},
						blkStmt([]ast.Node{
							expStmt(&ast.Identifier{
								ast.IDENTIFIER,
								"d",
							}),
						}),
						nil,
					},
				}),
			},
		}),
	},
	{
		`foreach (things as thing) { a; }`,
		blkStmt([]ast.Node{
			&ast.ForeachStatement{
				ast.FOREACH_STATEMENT,
				&ast.Identifier{
					ast.IDENTIFIER,
					"things",
				},
				nil,
				&ast.Identifier{
					ast.IDENTIFIER,
					"thing",
				},
				blkStmt([]ast.Node{
					expStmt(&ast.Identifier{
						ast.IDENTIFIER,
						"a",
					}),
				}),
			},
		}),
	},
	{
		`foreach (things as key=>thing) { a; }`,
		blkStmt([]ast.Node{
			&ast.ForeachStatement{
				ast.FOREACH_STATEMENT,
				&ast.Identifier{
					ast.IDENTIFIER,
					"things",
				},
				&ast.Identifier{
					ast.IDENTIFIER,
					"key",
				},
				&ast.Identifier{
					ast.IDENTIFIER,
					"thing",
				},
				blkStmt([]ast.Node{
					expStmt(&ast.Identifier{
						ast.IDENTIFIER,
						"a",
					}),
				}),
			},
		}),
	},
}

func TestStatements(t *testing.T) {

	statements.test(t);
}

var createObject = testCases {
	{
		`vo = 'value-object'(a, b);`,
		blkStmt([]ast.Node{
			expStmt(&ast.Infix{
				ast.INFIX,
				&ast.Identifier{ast.IDENTIFIER, "vo"},
				"=",
				&ast.ObjectCreation{
					ast.OBJECT_CREATION,
					"value-object",
					[]ast.Expression{
						&ast.Identifier{ast.IDENTIFIER, "a"},
						&ast.Identifier{ast.IDENTIFIER, "b"},
					},
				},
			}),
		}),
	},
	{
		`vo = 'vo-with-no-args';`,
		blkStmt([]ast.Node{
			expStmt(&ast.Infix{
				ast.INFIX,
				&ast.Identifier{ast.IDENTIFIER, "vo"},
				"=",
				&ast.ObjectCreation{
					ast.OBJECT_CREATION,
					"vo-with-no-args",
					[]ast.Expression{},
				},
			}),
		}),
	},
	{
		`'vo-with-no-args';`,
		blkStmt([]ast.Node{
			expStmt(&ast.ObjectCreation{
				ast.OBJECT_CREATION,
				"vo-with-no-args",
				[]ast.Expression{},
			}),
		}),
	},
}

func TestCreateObject(t *testing.T) {

	createObject.test(t);
}

var handlerSpecificStatementsAndExpressions = testCases {
	{
		`assert invariant 'is-started';`,
		blkStmt([]ast.Node{
			&ast.AssertStatement{
				ast.ASSERT_STATEMENT,
				"",
				&ast.ObjectCreation{
					ast.OBJECT_CREATION,
					"is-started",
					[]ast.Expression{},
				},
			},
		}),
	},
	{
		`assert invariant not 'is-started';`,
		blkStmt([]ast.Node{
			&ast.AssertStatement{
				ast.ASSERT_STATEMENT,
				"not",
				&ast.ObjectCreation{
					ast.OBJECT_CREATION,
					"is-started",
					[]ast.Expression{},
				},
			},
		}),
	},
	{
		`revision = run query 'next-revision-number' (agency_id, quote_number);`,
		blkStmt([]ast.Node{
			expStmt(&ast.Infix{
				ast.INFIX,
				&ast.Identifier{ast.IDENTIFIER, "revision"},
				"=",
				&ast.RunQuery{
					ast.RUN_QUERY,
					&ast.ObjectCreation{
						ast.OBJECT_CREATION,
						"next-revision-number",
						[]ast.Expression{
							&ast.Identifier{ast.IDENTIFIER, "agency_id"},
							&ast.Identifier{ast.IDENTIFIER, "quote_number"},
						},
					},
				},
			}),
		}),
	},
	{
		`apply event 'started';`,
		blkStmt([]ast.Node{
			&ast.ApplyStatement{
				ast.APPLY_STATEMENT,
				&ast.ObjectCreation{
					ast.OBJECT_CREATION,
					"started",
					[]ast.Expression{},
				},
			},
		}),
	},
}

func TestHandlerSpecificStatementsAndExpressions(t *testing.T) {

	handlerSpecificStatementsAndExpressions.test(t);
}


var precedenceTests = []struct {
	statement string
	expected  string
}{
	{
		"-a * b;",
		"((-a) * b);",
	},
	{
		"(a);",
		"a;",
	},
	{
		"!-a;",
		"(!(-a));",
	},
	{
		"a + b + c;",
		"((a + b) + c);",
	},
	{
		"a + b - c;",
		"((a + b) - c);",
	},
	{
		"a * b * c;",
		"((a * b) * c);",
	},
	{
		"a * b / c;",
		"((a * b) / c);",
	},
	{
		"a + b / c;",
		"(a + (b / c));",
	},
	{
		"a + b * c + d / e - f;",
		"(((a + (b * c)) + (d / e)) - f);",
	},
	{
		"3 + 4 - -5 * 5;",
		"((3 + 4) - ((-5) * 5));",
	},
	{
		"5 > 4 == 3 < 4;",
		"((5 > 4) == (3 < 4));",
	},
	{
		"5 < 4 != 3 > 4;",
		"((5 < 4) != (3 > 4));",
	},
	{
		"3 + 4 * 5 == 3 * 1 + 4 * 5;",
		"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)));",
	},
	{
		"3 > 5 == false;",
		"((3 > 5) == false);",
	},
	{
		"1 + (2 + 3) + 4;",
		"((1 + (2 + 3)) + 4);",
	},
	{
		"(5 + 5) * 2;",
		"((5 + 5) * 2);",
	},
	{
		"2 / (5 + 5);",
		"(2 / (5 + 5));",
	},
	{
		"(5 + 5) * 2 * (5 + 5);",
		"(((5 + 5) * 2) * (5 + 5));",
	},
	{
		"-(5 + 5);",
		"(-(5 + 5));",
	},
	{
		"!(true == true);",
		"(!(true == true));",
	},
	{
		"a->b->c->d;",
		"(((a -> b) -> c) -> d);",
	},
	{
		"a->b->c = 34 - 1;",
		"(((a -> b) -> c) = (34 - 1));",
	},
	{
		"a->b->c();",
		"((a -> b) -> c)();",
	},
	{
		"a->b(1, c + d);",
		"(a -> b)(1, (c + d));",
	},
	{
		"a->b()->c->d();",
		"(((a -> b)() -> c) -> d)();",
	},
	{
		"a->b[1];",
		"((a -> b)[1]);",
	},
	{
		"a->b[1][2];",
		"(((a -> b)[1])[2]);",
	},
}

func TestPredence(t *testing.T) {

	for _, testCase := range precedenceTests {

		statementBlockStr := "{ "+testCase.statement+" }";

		p := parser.NewStatement(statementBlockStr);

		node, err := p.ParseBlockStatement()

		if (err != nil) {
			t.Error("Got error on '"+testCase.statement+"'");
			t.Error(err.Error());
		}

		if (node.String() != testCase.expected) {
			t.Error("Error in precedence")
			t.Error("Expected: "+testCase.expected)
			t.Error("Got: "+node.String())
		}
	}
}

var invalidStatements = []struct{
	statement string
}{
	{"a a a"},
	{`foreach (things key=>thing) { a; }`},
	{`foreach (`},
	{`if a == b`},
	{`if (a) { b;`},
	{`a->b->c(d,`},
	{`a->b->c(d, e`},
	{`a[`},
	{`a[b`},
	{`a[b,c`},
}

func TestInvalidStatements(t *testing.T) {

	for _, testCase := range invalidStatements {

		p := parser.NewStatement("{"+testCase.statement+"}");

		node, err := p.ParseBlockStatement()

		if (err == nil) {
			t.Error("Expected error for '"+testCase.statement+"', got "+node.String())
		}
	}
}
