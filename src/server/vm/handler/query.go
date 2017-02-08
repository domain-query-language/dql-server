package handler

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

type Query interface {
	Id() vm.Identifier
	String() string
}
