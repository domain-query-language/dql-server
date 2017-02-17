package command

import (
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/value"
)

var TypeRename = uuid.FromStringOrNil("fca11037-1cbc-4e46-a7e4-077ffb90abc6")

type Rename struct {

	Name value.Name
}

func (self Rename) TypeId() vm.Identifier {
	return TypeRename
}
