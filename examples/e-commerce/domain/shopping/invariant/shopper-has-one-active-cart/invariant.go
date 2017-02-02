package shopper_has_one_active_cart

import "github.com/domain-query-language/dql-server/src/server/domain/vm"

type Invariant struct {

	shopper_id vm.Identifier
}

func (self *Invariant) Assumptions(shopper_id vm.Identifier) {
	self.shopper_id = shopper_id
}

func (self *Invariant) Satisfier(queryable Queryable) bool {
	return !queryable.IsActive(self.shopper_id)
}
