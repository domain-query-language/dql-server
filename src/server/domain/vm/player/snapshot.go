package player

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"time"
)

type Snapshot struct {

	Id vm.Identifier

	LastId vm.Identifier

	OccurredAt time.Time

	Version int
}

func (self *Snapshot) Encode() []byte {
	return []byte{}
}

func (self *Snapshot) Decode([]byte) {

}
