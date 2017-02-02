package event

import (
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

var TypeCreated, _ = uuid.FromString("82e81f55-2780-4685-894e-498b94607f00")

type Created struct {

	ShopperId vm.Identifier
}

