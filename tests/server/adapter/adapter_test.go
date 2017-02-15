package adapter

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/domain-query-language/dql-server/examples/dql/application/projection/list-databases"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/database"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/database/command"
	"github.com/domain-query-language/dql-server/src/server/adapter"
	"github.com/domain-query-language/dql-server/src/server/adapter/parser"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/satori/go.uuid"
	"testing"
)

var listStatements = []struct{
	statement string
	expected *adapter.Handleable
}{
	{
		"LIST DATABASES;",
		adapter.NewQuery( vm.NewQuery(
			list_databases.Identifier,
			list_databases.Query{},
		), ),
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

		} else if (testCase.expected.Query.Id() != actual.Query.Id()) {
			t.Error("Expected query cases to match");
			t.Error("Expected: "+testCase.expected.Query.String())
			t.Error("Got: "+actual.Query.String())
		}
	}
}

var id, _ = uuid.FromString("2bdecde9-a3a2-43cd-a1b6-234855e5399a")

var createStatements = []struct{
	statement string
	expected *adapter.Handleable
}{
	{
		"CREATE DATABASE 'db';",
		adapter.NewCommand(
			vm.NewCommand(
				vm.NewAggregateIdentifier(id, database.Identifier),
				command.Create {"db"},
			),

		),
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

		same, errMsg := commandsAreSame(actual.Command, testCase.expected.Command)

		if (!same) {

			expectedAsStr := spew.Sdump(testCase.expected)
			actualAsStr := spew.Sdump(actual)

			t.Error("Commands do not match: "+errMsg);
			t.Error("Expected: "+expectedAsStr)
			t.Error("Got: "+actualAsStr)
		}
	}
}

func commandsAreSame(actual vm.Command, expected vm.Command) (result bool, error string) {

	if (actual == nil || expected == nil) {
		return false, "Command objects are missing"
	}

	if (actual.TypeId() != expected.TypeId()) {
		return false, "Types of commands do not match."
	}

	if (string(actual.AggregateId().Bytes()) != string(expected.AggregateId().Bytes())) {
		return false, "Not referencing the same aggregate ID/Type"
	}

	if (actual.Payload() != expected.Payload()) {
		return false, "Command Payloads do not match"
	}

	return true, ""
}

var invalidStatements = []struct{
	statement string
	error string
}{
	{
		"LIST DATABASES",
		"Error at char 14, expected [;], got [eof] instead",
	},
	{
		"LIST BANANAS;",
		"Error at char 5, expected 'databases', got 'BANANAS' instead",
	},
	{
		"CREATE DATABASE tim;",
		"Error at char 16, expected [objectName], got [identifier] instead",
	},
	{
		"DATABASE",
		"Error at char 0, expected [create/list], got [identifier] instead",
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
			t.Error("Expected error: "+testCase.error);
		} else if (err.Error() != testCase.error) {
			t.Error("Got error on '"+testCase.statement+"'")
			t.Error("Errors do not match");
			t.Error("Expected: "+testCase.error);
			t.Error("Got: "+err.Error());
		}
	}
}

