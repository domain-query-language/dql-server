package database_name_unique

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/projection/database-name-unique"
)

var Projector = projection.NewProjector(
	Projection,
	database_name_unique.ProjectorHandlers,
)

var Identifier = database_name_unique.Identifier
