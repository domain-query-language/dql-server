package store
/*
import (
	"github.com/satori/go.uuid"
	"github.com/boltdb/bolt"
	"bytes"
	"log"
	"encoding/binary"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

type BoltStream_ struct {

	current_id uuid.UUID
	current_event_key int

	current_event vm.Loggable

	db *bolt.DB
}

func (self *BoltStream_) Reset() {

	self.current_id = uuid.Nil
	self.current_event_key = 0
}

func (self *BoltStream_) LastId() vm.Identifier {
	return self.current_id
}

func (self *BoltStream_) Seek(identifier vm.Identifier) {
	panic("Seek not implemented for BoltStream.")
}

func (self *BoltStream_) Next() bool {

	start := binary.BigEndian.Uint64(self.current_event_key)

	self.db.View(func(tx *bolt.Tx) error {

		events_log := tx.Bucket([]byte("events_log"))
		cursor := events_log.Cursor()

		for k, v := cursor.Seek(self.event_key); k != nil && bytes.Compare(k, max) <= 0; k, v = cursor.Next() {

			event := &sourced.Event_{}
			event.GobDecode(v)

			events = append(events, event)

			self.last_id = event.Id()
			self.event_key = k
			self.stream_count++
		}

		return nil
	})

	return events
}

func (self *BoltStream_) Value() vm.Loggable {
	return self.current_event
}

func (self *BoltStream_) Close() {
	self.db.Close()
}

func BoltStream(path string) *BoltStream_ {

	db, err := bolt.Open(
		path,
		0666,
		&bolt.Options{ReadOnly: true},
	)

	if err != nil {
		log.Fatal(err)
	}

	bolt_stream := &BoltStream_ {
		current_id: uuid.Nil,
		current_event_key: 0,
		db: db,
	}

	bolt_stream.Reset()

	return bolt_stream
}
*/