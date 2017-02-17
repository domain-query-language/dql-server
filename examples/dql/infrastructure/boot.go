package infrastructure

import (
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/aggregate/database"
	"github.com/domain-query-language/dql-server/examples/dql/infrastructure/application/projection/list-databases"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/player"
	"github.com/domain-query-language/dql-server/examples/dql/application"
	"github.com/domain-query-language/dql-server/examples/dql/infrastructure/domain/projection/database-name-unique"
)

func Boot() {

	AggregatesRepository.Add(database.Aggregate)

	ProjectionsRepository.Add(list_databases.Projection)

	PlayersRepository.Add(
		player.NewPlayer(
			database_name_unique.Identifier,
			modelling.Identifier,
			EventLog.Stream(),
			database_name_unique.Projector,
		),
	)

	PlayersRepository.Add(
		player.NewPlayer(
			list_databases.Identifier,
			application.Identifier,
			EventLog.Stream(),
			list_databases.Projector,
		),
	)

	QueryHandler.Add(
		list_databases.Identifier,
		list_databases.QueryHandler,
	)
}
