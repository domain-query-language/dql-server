package store

import "time"

type Event interface {

	Id() Identifier

	TypeId() Identifier

	OccurredAt() time.Time

	GobEncode() []byte

	GobDecode([]byte) error
}

type Event_ struct {

	id Identifier
	type_id Identifier

	occurred_at time.Time

	payload []byte
}
