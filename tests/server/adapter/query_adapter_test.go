package adapter

import (
	"testing"
	"github.com/domain-query-language/dql-server/src/server/adapter"
	schemaAdapter "github.com/domain-query-language/dql-server/src/server/adapter/schema"
	query "github.com/domain-query-language/dql-server/src/server/query/schema"
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

		adptr := schemaAdapter.NewQueryAdapter(testCase.statement);

		actual, err := adptr.Next();

		if (err != nil) {
			t.Error("Got error")
			t.Error(err);
		}

		if (actual.Query == nil) {

			t.Error("Query cannot be nil, expected valid query object")

		} else if (testCase.expected.Query != actual.Query) {

			t.Error("Expected query cases to match");
			t.Error("Expected: "+testCase.expected.Query.String())
			t.Error("Got: "+actual.Query.String())
		}
	}
}
