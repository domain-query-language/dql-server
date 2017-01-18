package projector

import (
	"github.com/domain-query-language/dql-server/src/server/vm"
	"time"
)

type Snapshot interface {

	Id() vm.Identifier

	OccurredAt() time.Time

	Version() int

	Payload() []byte

	GobEncode() []byte
}

type Snapshot_ struct {

	id vm.Identifier

	occurred_at time.Time

	version int

	payload []byte
}

func CreateSnapshot(id vm.Identifier, version int, payload []byte) Snapshot {

	return &Snapshot_ {
		id: id,
		version: version,
		occurred_at: time.Now(),
		payload: payload,
	}
}
