package shopper_has_one_active_cart

import "github.com/domain-query-language/dql-server/src/server/domain/vm"

type Queryable interface {

	IsActive(shopper_id vm.Identifier) bool
}
