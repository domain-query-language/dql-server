package adapter

import (
	"github.com/domain-query-language/dql-server/src/server/adapter"
	"strings"
	"errors"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"regexp"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/aggregate/database/command"
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/examples/dql/application/projection/list-databases"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/aggregate/database"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/value"
)

var CREATE_DATABASE_REGEX = regexp.MustCompile("^create database \\'([a-zA-Z0-9-]{1,256})\\'$")

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
				list_databases.Identifier,
				list_databases.Query{

				},
			),
		), nil
	}

	if CREATE_DATABASE_REGEX.MatchString(statement) {

		self.index++

		name, name_err := value.NewName(
			CREATE_DATABASE_REGEX.FindStringSubmatch(statement)[1],
		)

		if name_err != nil {
			return nil, name_err
		}

		return adapter.NewCommand(
			vm.NewCommand(
				vm.NewAggregateIdentifier(
					uuid.NewV4(),
					database.Identifier,
				),
				command.Create {
					Name: name,
				},
			),
		), nil
	}

	return nil, errors.New("The command/query given does not exist.")
}

func NewMockAdapter(input string) *MockAdapter {

	return &MockAdapter {
		statements: strings.Split(input, ";"),
		index: 0,
	}
}
