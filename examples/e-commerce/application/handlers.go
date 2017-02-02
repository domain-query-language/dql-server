package application

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/handler/command"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

var CommandHandler = command.NewSimpleHandler(
	map[vm.Identifier]vm.Identifier{},
	AggregateRepository,
	PlayerRepository,
)
