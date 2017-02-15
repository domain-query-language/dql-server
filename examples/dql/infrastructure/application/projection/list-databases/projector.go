package list_databases

import (
	"github.com/domain-query-language/dql-server/examples/dql/application/projection/list-databases"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
)

var Projector = projection.NewProjector(
	Projection,
	list_databases.ProjectorHandlers,
)

var Identifier = list_databases.Identifier

var QueryHandler = list_databases.QueryHandler