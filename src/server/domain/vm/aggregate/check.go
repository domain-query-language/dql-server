package aggregate

import "github.com/domain-query-language/dql-server/src/server/domain/vm"

type Check interface {

	That(invariant_id vm.Identifier) Check


}
