package application

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/handler/query"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/handler/command"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/database"
)

var QueryHandler = query.CreateHandler(ProjectionsRepository)

var CommandHandler = command.NewHandler(
	map[vm.Identifier]vm.Identifier {
		database.Identifier: database.Aggregate,
	},
	AggregatesRepository,
	PlayersRepository,
)
