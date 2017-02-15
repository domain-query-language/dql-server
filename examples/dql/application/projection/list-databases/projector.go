package list_databases

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/database/event"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
	"github.com/satori/go.uuid"
)

var Identifier = uuid.FromStringOrNil("c50d6791-3fc5-4be8-91fc-c01f20526872")

var ProjectorHandlers = &map[vm.Identifier]projection.ProjectorHandler {

	event.TypeCreated: func(p projection.Projection, e vm.Event) {

		projection := p.(Projection)
		event := e.Payload().(event.Created)

		projection.Add(event.Name)
	},
}
