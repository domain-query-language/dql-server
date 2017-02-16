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

	command.TypeCreate: func(agg aggregate.Aggregate, c vm.Command) error  {

		create := c.Payload().(command.Create)

		//agg.Check().That(invariant.Created).Not().Asserts()

		agg.Apply(
			vm.NewEvent(
				*c.AggregateId(),
				c.Id(),
				event.Created {
					create.Name,
				},
			),
		)

		return nil
	},

	command.TypeRename: func(agg aggregate.Aggregate, c vm.Command) error  {

		rename := c.Payload().(command.Rename)

		// Assert Invariant 'Created'

		agg.Apply(
			vm.NewEvent(
				*c.AggregateId(),
				c.Id(),
				event.Renamed {
					rename.Name,
				},
			),
		)

		return nil
	},

	command.TypeDelete: func(agg aggregate.Aggregate, c vm.Command) error  {

		// Assert Invariant 'Created'

		agg.Apply(
			vm.NewEvent(
				*c.AggregateId(),
				c.Id(),
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
