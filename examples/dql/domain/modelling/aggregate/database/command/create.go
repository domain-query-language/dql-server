package command

import (
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/value"
)

var TypeCreate = uuid.FromStringOrNil("b453fd7a-dd41-4b09-9d3f-c765027b1bf4")

type Create struct {

	Name value.Name
}

func (self Create) TypeId() vm.Identifier {
	return TypeCreate
}
