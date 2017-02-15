package aggregate

import (
	"time"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

type Snapshot struct {

	Id vm.Identifier

	OccurredAt time.Time

	Version int

	Payload []byte
}

func (self *Snapshot) Encode() []byte {
	return []byte{}
}
