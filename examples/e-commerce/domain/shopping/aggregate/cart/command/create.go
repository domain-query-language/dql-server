package command

import (
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

var TypeCreate, _ = uuid.FromString("ef12a076-c466-49e2-a062-615eacb94bcd")

type Create struct {

	ShopperId vm.Identifier
}
