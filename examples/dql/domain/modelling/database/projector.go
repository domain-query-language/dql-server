package database

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/database/event"
	"github.com/davecgh/go-spew/spew"
)

var ProjectorHandlers = &map[vm.Identifier]projection.ProjectorHandler {

	event.TypeCreated: func(p projection.Projection, e vm.Event) {

		projection := p.(Projection)

		projection.Create()

		spew.Dump(projection)
	},
}

var Projector = projection.NewProjector(
	State,
	ProjectorHandlers,
)
