package adapter

import (

	"errors"
	"github.com/domain-query-language/dql-server/src/server/adapter"
	"github.com/domain-query-language/dql-server/src/server/adapter/parser"
	query "github.com/domain-query-language/dql-server/src/server/query/schema"
	command "github.com/domain-query-language/dql-server/src/server/command/schema"
	"testing"
	"github.com/davecgh/go-spew/spew"
)

var listStatements = []struct{
	statement string
	expected *adapter.Handleable
}{
	{
		"LIST DATABASES;",
		adapter.NewQuery( &query.ListDatabases{} ),
	},
}

func TestStatementToListQuery(t *testing.T){

	for _, testCase := range listStatements {

		adptr := parser.New(testCase.statement);

		actual, err := adptr.Next();

		if (err != nil) {

			t.Error("Got error on '"+testCase.statement+"'")
			t.Error(err);
		}

		if (actual == nil) {

			t.Error("Query cannot be nil, expected valid query object")

		} else if (testCase.expected.Query != actual.Query) {
			t.Error("Expected query cases to match");
			t.Error("Expected: "+testCase.expected.Query.String())
			t.Error("Got: "+actual.Query.String())
		}
	}
}

var createStatements = []struct{
	statement string
	expected *adapter.Handleable
}{
	{
		"CREATE DATABASE 'db';",
		adapter.NewCommand(&command.CreateDatabase{"db"}),
	},
}

func TestCreateStatements(t *testing.T){

	spew.Config.DisableMethods = true;
	spew.Config.DisablePointerAddresses = true;

	for _, testCase := range createStatements {

		adptr := parser.New(testCase.statement);

		actual, err := adptr.Next();

		if (err != nil) {

			t.Error("Got error on '"+testCase.statement+"'")
			t.Error(err);
		}

		expectedAsStr := spew.Sdump(testCase.expected)
		actualAsStr := spew.Sdump(actual)

		if (actual == nil) {

			t.Error("Command cannot be nil, expected valid command object")

		} else if (expectedAsStr != actualAsStr) {
			t.Error("Expected query cases to match");
			t.Error("Expected: "+expectedAsStr)
			t.Error("Got: "+actualAsStr)
		}
	}
}

var invalidStatements = []struct{
	statement string
	error error
}{
	{
		"LIST DATABASES",
		errors.New("Expected next token to be ';', got EOF instead"),
	},
	{
		"LIST BANANAS;",
		errors.New("Expected next token to be 'databases', got 'Parse error, expected databases, found BANANAS' instead"),
	},
	{
		"CREATE BANANAS;",
		errors.New("Expected next token to be 'database', got 'Parse error, expected database, domain, context, aggregate, found BANANAS' instead"),
	},
}

func TestInvalidStatement(t *testing.T) {

	for _, testCase := range invalidStatements {

		adptr := parser.New(testCase.statement);

		actual, err := adptr.Next();

		if (actual != nil) {
			t.Error("Got error on '"+testCase.statement+"'")
			t.Error("Expected error, got object");
			t.Error("Got object: "+actual.String());
			t.Error("Expected error: "+testCase.error.Error());
		} else if (err.Error() != testCase.error.Error()) {
			t.Error("Got error on '"+testCase.statement+"'")
			t.Error("Errors do not match");
			t.Error("Expected: "+testCase.error.Error());
			t.Error("Got: "+err.Error());
		}
	}
}

