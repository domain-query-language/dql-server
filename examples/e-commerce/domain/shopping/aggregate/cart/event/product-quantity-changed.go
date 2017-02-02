package event

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/examples/e-commerce/domain/shopping/value"
	"github.com/satori/go.uuid"
)

var TypeProductQuantityChanged, _ = uuid.FromString("7f0e0c38-8138-4bad-a916-56428935f337")

type ProductQuantityChanged struct {

	ProductId vm.Identifier
	Quantity value.Quantity
}
