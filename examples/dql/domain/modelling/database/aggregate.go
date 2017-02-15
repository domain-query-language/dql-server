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
				cmd.Id(),
				event.Created {
					payload.Name,
				},
			),
		)

		return nil
	},

	command.TypeRename: func(aggregate aggregate.Aggregate, cmd vm.Command) error  {

		payload := cmd.Payload().(command.Rename)

		// Assert Invariant 'Created'

		aggregate.Apply(
			vm.NewEvent(
				cmd.AggregateId(),
				cmd.Id(),
				event.Renamed {
					payload.Name,
				},
			),
		)

		return nil
	},

	command.TypeDelete: func(aggregate aggregate.Aggregate, cmd vm.Command) error  {

		// Assert Invariant 'Created'

		aggregate.Apply(
			vm.NewEvent(
				cmd.AggregateId(),
				cmd.Id(),
				event.Deleted {},
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
