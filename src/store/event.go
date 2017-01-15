package store

import "time"

type Event interface {

	Id() Identifier

	TypeId() Identifier

	AggregateId() Identifier

	AggregateTypeId() Identifier

	OccurredAt() time.Time

	GobEncode() []byte

	GobDecode([]byte) error
}
