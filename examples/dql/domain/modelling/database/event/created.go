package event

import (
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/valueobjects"
)

var TypeCreated = uuid.FromStringOrNil("b7dbb816-1141-4357-abbf-250e0cb7ec1f")

type Created struct {

	Name valueobjects.Name
}

func (self Created) TypeId() vm.Identifier {
	return TypeCreated
}