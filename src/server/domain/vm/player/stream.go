package player

import "github.com/domain-query-language/dql-server/src/server/domain/vm"

type Stream interface {

	LastId() vm.Identifier

	Reset()

	Seek(id vm.Identifier) error

	Next() (vm.Event, error)
}
