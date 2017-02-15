package store

import "github.com/domain-query-language/dql-server/src/server/domain/vm"

type Stream interface {

	Reset()

	LastId() vm.Identifier

	Seek(identifier vm.Identifier)

	Next() bool

	Value() vm.Event

}
