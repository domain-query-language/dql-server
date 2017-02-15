package list_databases

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/database/event"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
)

var ProjectorHandlers = &map[vm.Identifier]projection.ProjectorHandler {

	event.TypeCreated: func(p projection.Projection, e vm.Event) {

		projection := p.(Projection)
		event := e.Payload().(event.Created)

		projection.Add(event.Name)
	},
}

