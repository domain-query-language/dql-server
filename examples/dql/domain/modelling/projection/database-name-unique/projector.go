package database_name_unique

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/aggregate/database/event"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
	"github.com/satori/go.uuid"
)

var Identifier = uuid.FromStringOrNil("491a2acf-6c58-45e0-b151-1f5ca306905f")

var ProjectorHandlers = &map[vm.Identifier]projection.ProjectorHandler {

	event.TypeCreated: func(p projection.Projection, e vm.Event) {

		projection := p.(Projection)
		event := e.Payload().(event.Created)

		projection.Add(e.AggregateId().Id, string(event.Name))
	},

	event.TypeRenamed: func(p projection.Projection, e vm.Event) {

		projection := p.(Projection)
		event := e.Payload().(event.Renamed)

		projection.Rename(e.AggregateId().Id, string(event.Name))
	},

	event.TypeDeleted: func(p projection.Projection, e vm.Event) {

		projection := p.(Projection)

		projection.Remove(e.AggregateId().Id)
	},

}
