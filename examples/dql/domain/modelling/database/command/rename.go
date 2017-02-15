package command

import (
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

var TypeRename = uuid.FromStringOrNil("fca11037-1cbc-4e46-a7e4-077ffb90abc6")

type Rename struct {

	Name string
}

func (self Rename) TypeId() vm.Identifier {
	return TypeRename
}
