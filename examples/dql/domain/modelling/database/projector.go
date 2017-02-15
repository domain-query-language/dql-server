package database

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/database/event"
)

var ProjectorHandlers = &map[vm.Identifier]projection.ProjectorHandler {

	event.TypeCreated: func(projection Projection, event vm.Event) {

		projection.Create()
	},
}

var Projector = projection.NewProjector(
	Projection,
	ProjectorHandlers,
)
