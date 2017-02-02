package cart

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/examples/e-commerce/domain/shopping/aggregate/cart/event"
)

var AggregateProjection = NewProjection(AggregateIdentifier)

var AggregateProjector = projection.NewProjector(AggregateProjection,

	map[vm.Identifier]projection.ProjectorHandler {

		event.TypeCreated : func(projection Projection, event vm.Event) {

			projection.Create()
		},

		event.TypeProductAdded : func(projection Projection, event vm.Event) {

			projection.AddProduct(
				event.(event.ProductAdded).Product,
			)
		},

		event.TypeProductQuantityChanged : func(projection Projection, event vm.Event) {

			e := event.(event.ProductQuantityChanged)

			projection.ChangeProductQuantity(
				e.ProductId,
				e.Quantity,
			)
		},

		event.TypeProductRemoved : func(projection Projection, event vm.Event) {

			projection.RemoveProduct(
				event.(event.ProductRemoved).ProductId,
			)
		},

		event.TypeCheckedOut : func(projection Projection, event vm.Event) {

			projection.Checkout()
		},

	},
)
