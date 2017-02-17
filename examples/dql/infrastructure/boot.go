package infrastructure

import (
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/database"
	"github.com/domain-query-language/dql-server/examples/dql/infrastructure/application/projection/list-databases"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/player"
)

func Boot() {

	AggregatesRepository.Add(database.Aggregate)

	ProjectionsRepository.Add(list_databases.Projection)

	PlayersRepository.Add(
		player.NewPlayer(
			list_databases.Identifier,
			modelling.Identifier,
			EventLog.Stream(),
			list_databases.Projector,
		),
	)

	QueryHandler.Add(
		list_databases.Identifier,
		list_databases.QueryHandler,
	)
}

func Reset() {

	list_databases.Projection.Reset()
}
