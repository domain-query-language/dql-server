package sourced

import (
	"github.com/satori/go.uuid"
	"bytes"
	"encoding/gob"
	"time"
	"github.com/lyonscf/boltdb-test/src/store"
)

type Event_ struct {

	id uuid.UUID
	type_id uuid.UUID
	aggregate_id uuid.UUID
	aggregate_type_id uuid.UUID

	occurred_at time.Time

	version int8

	payload []byte
}

func (self *Event_) Id() store.Identifier {
	return self.id
}

func (self *Event_) TypeId() store.Identifier {
	return self.type_id
}

func (self *Event_) AggregateId() store.Identifier {
	return self.aggregate_id
}

func (self *Event_) AggregateTypeId() store.Identifier {
	return self.aggregate_type_id
}

func (self *Event_) OccurredAt() time.Time {
	return self.occurred_at
}

func (self *Event_) GobEncode() []byte {

	buffer := new(bytes.Buffer)

	encCache := gob.NewEncoder(buffer)
	encCache.Encode(self)

	return buffer.Bytes()
}

func (self *Event_) GobDecode([]byte) error {

	buffer := new(bytes.Buffer)

	encCache := gob.NewDecoder(buffer)
	return encCache.Decode(self)
}

func Event(
	type_id uuid.UUID,
	aggregate_id uuid.UUID,
	aggregate_type_id uuid.UUID,
	version int8,
	payload []byte,
) store.Event {

	return &Event_ {
		id: uuid.NewV4(),
		type_id: type_id,
		aggregate_id: aggregate_id,
		aggregate_type_id: aggregate_type_id,
		occurred_at: time.Now(),
		version: version,
		payload: payload,
	}
}
