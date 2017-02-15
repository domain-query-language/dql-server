package application

import (
	"github.com/domain-query-language/dql-server/src/server/infrastructure/domain/vm/projection"
	"github.com/domain-query-language/dql-server/src/server/infrastructure/domain/vm/player"
	"github.com/domain-query-language/dql-server/src/server/infrastructure/domain/vm/aggregate"
)

var ProjectionsRepository = projection.CreateMemoryRepository()

var PlayersRepository = player.CreateMemoryRepository()

var AggregatesRepository = aggregate.CreateMemoryRepository()
