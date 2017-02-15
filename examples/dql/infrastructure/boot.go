package infrastructure

import (
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/database"
	"github.com/domain-query-language/dql-server/examples/dql/infrastructure/application/projection"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/player"
	"github.com/domain-query-language/dql-server/examples/dql/application/projection/list-databases"
)

func Boot() {

	ProjectionsRepository.Save(
		projection.ListDatabasesProjection,
	)

	QueryHandler.Add(
		list_databases.Identifier,
		list_databases.QueryHandler,
	)

	AggregatesRepository.Add(
		database.Identifier,
		database.Aggregate,
	)

	PlayersRepository.Add(
		player.NewPlayer(
			list_databases.Identifier,
			modelling.Identifier,
			EventLog.Stream(),
			projection.ListDatabasesProjector,
		),
	)
}
