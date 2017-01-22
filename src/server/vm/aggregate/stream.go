package aggregate

import "github.com/domain-query-language/dql-server/src/server/vm"

type Stream interface {

	Reset()

	Seek(version int)

	Next() vm.Event
}
