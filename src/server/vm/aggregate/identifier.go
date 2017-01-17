package aggregate

import (
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/src/server/vm"
)

type Identifier struct {

	id uuid.UUID
	typeId uuid.UUID

}

func (self *Identifier) Bytes() []byte {
	return append(self.id.Bytes(), self.typeId.Bytes()...)
}

func CreateIdentifier(id vm.Identifier, type_id vm.Identifier) *Identifier {

	return &Identifier{
		id: id,
		typeId: type_id,
	}
}
