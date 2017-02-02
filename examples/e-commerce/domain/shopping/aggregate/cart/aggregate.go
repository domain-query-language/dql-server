package cart

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/aggregate"
	"github.com/domain-query-language/dql-server/examples/e-commerce/domain/shopping/aggregate/cart/command"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/examples/e-commerce/domain/shopping/aggregate/cart/event"
)

var AggregateIdentifier, _ = uuid.FromString("7a742e09-3503-4c56-b1a5-edea571a8227")

var Aggregate = aggregate.CreateAggregate(AggregateIdentifier, AggregateProjector,

	map[vm.Identifier]aggregate.AggregateHandler {

		command.TypeCreate: func(projector projection.Projector, command vm.Command) ([]vm.Event, error)  {

			cmd := command.(command.Create)

			// Assert Invariant 'Created'

			// Assert Invariant 'ShopperHasOneActiveCart'

			projector.Apply(
				vm.NewEvent(
					event.TypeCreated,
					command.Id(),
					command.AggregateId(),
					command.ContextId(),
					event.Created {
						ShopperId: cmd.ShopperId,
					},
				),
			)

			projector.Apply(
				vm.NewEvent(
					event.TypeEmpty,
					command.Id(),
					command.AggregateId(),
					command.ContextId(),
					event.Empty {},
				),
			)

			return nil, nil
		},
	},
)
