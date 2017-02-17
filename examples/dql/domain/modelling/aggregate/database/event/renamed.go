package event

import (
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/value"
)

var TypeRenamed = uuid.FromStringOrNil("49f1e1b1-f319-4f41-ad22-f6061f3476a7")

type Renamed struct {

	Name value.Name
}

func (self Renamed) TypeId() vm.Identifier {
	return TypeRenamed
}
