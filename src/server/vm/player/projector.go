package player

import "github.com/domain-query-language/dql-server/src/server/vm"

type Projector interface {

	Reset()

	Apply(event vm.Event) error

	Version() int
}
