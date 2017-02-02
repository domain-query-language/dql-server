package shopper_has_one_active_cart

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

type Projection interface {

	projection.Projection

	Create(cart_id vm.Identifier)

	Checkout(cart_id vm.Identifier)
}
