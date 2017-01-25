package projection

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"time"
)

type Snapshot interface {

	Id() vm.Identifier

	OccurredAt() time.Time

	Payload() []byte

	GobEncode() []byte
}

type Snapshot_ struct {

	id vm.Identifier

	occurred_at time.Time

	payload []byte
}

func (self *Snapshot_) Id() vm.Identifier {
	return self.id
}

func (self *Snapshot_) OccurredAt() time.Time {
	return self.occurred_at
}

func (self *Snapshot_) Payload() []byte {
	return self.payload
}

func (self *Snapshot_) GobEncode() []byte {
	return []byte{}
}

func CreateSnapshot(id vm.Identifier, payload []byte) Snapshot {

	return &Snapshot_ {
		id: id,
		occurred_at: time.Now(),
		payload: payload,
	}
}
