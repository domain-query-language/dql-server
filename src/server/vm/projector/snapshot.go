package projector

import (
	"github.com/domain-query-language/dql-server/src/server/vm"
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

func CreateSnapshot(id vm.Identifier, payload []byte) Snapshot {

	return &Snapshot_ {
		id: id,
		occurred_at: time.Now(),
		payload: payload,
	}
}
