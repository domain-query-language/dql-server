package entity

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/examples/e-commerce/domain/shopping/value"
)

type Product struct {

	Id vm.Identifier
	Quantity value.Quantity
}
