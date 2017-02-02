package application

import (
	"github.com/domain-query-language/dql-server/src/server/infrastructure/domain/vm/player"
	"github.com/domain-query-language/dql-server/src/server/infrastructure/domain/vm/projection"
)

var ProjectionRepository = projection.CreateMemoryRepository()

var AggregateRepository = projection.CreateMemoryRepository()

var PlayerRepository = player.CreateMemoryRepository()
