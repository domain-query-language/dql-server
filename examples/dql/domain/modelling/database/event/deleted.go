package event

import (
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

var TypeDeleted = uuid.FromStringOrNil("f4babc1d-8d6b-47ae-9a3f-bb3b2bccc1ff")

type Deleted struct {

}

func (self Deleted) TypeId() vm.Identifier {
	return TypeDeleted
}
