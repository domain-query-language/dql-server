package database

import (
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/aggregate"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/database/command"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/database/event"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling"
)

var Identifier = uuid.FromStringOrNil("655bdd1e-8deb-4e08-9bbf-14496e148e2e")

var Handlers = 	&map[vm.Identifier]aggregate.AggregateHandler {

	command.TypeCreate: func(aggregate aggregate.Aggregate, cmd vm.Command) error  {

		payload := cmd.Payload().(command.Create)

		// Assert Invariant 'Created'

		aggregate.Apply(
			vm.NewEvent(
				cmd.AggregateId(),
				event.Created {
					payload.Name,
				},
			),
		)

		return nil
	},
}

var Aggregate = aggregate.NewAggregate(
	Identifier,
	modelling.Identifier,
	Projector,
	Handlers,
)
