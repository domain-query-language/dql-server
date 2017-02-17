package command

import (
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

var TypeDelete = uuid.FromStringOrNil("60dbb5f2-61b3-40e2-b606-0c4ab09a5bcf")

type Delete struct {

}

func (self Delete) TypeId() vm.Identifier {
	return TypeDelete
}
