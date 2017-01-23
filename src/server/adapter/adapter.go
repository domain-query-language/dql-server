package adapter

import (
	"github.com/domain-query-language/dql-server/src/server/vm/handler/command"
	"github.com/domain-query-language/dql-server/src/server/vm/handler/query"
)

type CommandAdapter interface {
	Next() (command.Command, error)
}

type QueryAdapter interface {
	Next() (query.Query, error)
}

