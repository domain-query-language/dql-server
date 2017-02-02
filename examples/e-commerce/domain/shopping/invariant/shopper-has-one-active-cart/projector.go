package shopper_has_one_active_cart

import (
	"github.com/domain-query-language/dql-server/examples/e-commerce/domain/shopping/aggregate/cart/event"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
	"github.com/domain-query-language/dql-server/examples/e-commerce/application"
	"github.com/satori/go.uuid"
)

var PlayerId, _ = uuid.FromString("062c2616-b7e6-4303-883c-1b083527195b")

var Projector = projection.NewProjector(application.ProjectionRepository.Get(PlayerId),

	map[vm.Identifier]projection.ProjectorHandler {

		event.TypeCreated : func(projection Projection, event vm.Event) {

			projection.Create(
				event.Id(),
			)
		},

		event.TypeCheckedOut : func(projection Projection, event vm.Event) {

			projection.Checkout(
				event.Id(),
			)
		},

	},
)
