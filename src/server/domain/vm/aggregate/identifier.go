package aggregate

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

type Identifier struct {

	id vm.Identifier
	typeId vm.Identifier

}

func (self *Identifier) Id() vm.Identifier {
	return self.id
}

func (self *Identifier) TypeId() vm.Identifier {
	return self.typeId
}

func (self *Identifier) Bytes() []byte {
	return append(self.id.Bytes(), self.typeId.Bytes()...)
}

func CreateIdentifier(id vm.Identifier, type_id vm.Identifier) *Identifier {

	return &Identifier {
		id: id,
		typeId: type_id,
	}
}
