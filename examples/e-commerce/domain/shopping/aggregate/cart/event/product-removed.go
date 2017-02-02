package event

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/satori/go.uuid"
)

var TypeProductRemoved, _ = uuid.FromString("8511d45e-6a0f-48c7-b6f1-d9521455dc7c")

type ProductRemoved struct {

	ProductId vm.Identifier
}
