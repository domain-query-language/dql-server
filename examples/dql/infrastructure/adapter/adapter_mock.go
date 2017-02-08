package adapter

import (
	"github.com/domain-query-language/dql-server/src/server/adapter"
	"strings"
	"errors"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/examples/dql/application/query/list-databases"
)

type MockAdapter struct {

	statements []string
	index int
}

func (self *MockAdapter) Next() (*adapter.Handleable, error) {

	statement := self.statements[self.index]
	statement = strings.TrimSpace(statement)

	if statement == "" {
		return nil, errors.New("There are no statements to process.")
	}

	if(statement == "list databases") {

		self.index++

		return adapter.NewQuery(
			vm.NewQuery(
				list_databases.QueryId,
				list_databases.ListDatabases{},
			),
		), nil
	}

	return nil, errors.New("The query given does not exist.")
}

func NewMockAdapter(input string) *MockAdapter {

	return &MockAdapter {
		statements: strings.Split(input, ";"),
		index: 0,
	}
}
