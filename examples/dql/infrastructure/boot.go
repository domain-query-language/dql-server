package infrastructure

import (
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/database"
	"github.com/domain-query-language/dql-server/examples/dql/infrastructure/application"
	"github.com/domain-query-language/dql-server/examples/dql/infrastructure/application/projection"
	"github.com/domain-query-language/dql-server/examples/dql/application/query/list-databases"
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/player"
)

func Boot() {

	application.ProjectionsRepository.Save(
		projection.ListDatabasesProjection,
	)

	application.QueryHandler.Add(
		projection.ListDatabasesProjectionID,
		list_databases.Handler,
	)

	application.AggregatesRepository.Add(
		database.Identifier,
		database.Aggregate,
	)

	application.PlayersRepository.Add(
		player.NewPlayer(
			uuid.FromStringOrNil("c50d6791-3fc5-4be8-91fc-c01f20526872"),
			modelling.Identifier,
			application.EventLog.Stream(),
			projection.ListDatabasesProjector,
		),
	)
}
