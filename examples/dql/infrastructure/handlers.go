package infrastructure

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/handler/query"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/handler/command"
)

var QueryHandler = query.CreateHandler(ProjectionsRepository)

var CommandHandler = command.NewHandler(
	AggregatesRepository,
	PlayersRepository,
)
