package store

/*
import (
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/sourced"
	"github.com/boltdb/bolt"
	"bytes"
	"log"
	"encoding/binary"
)

type BoltStream_ struct {

	starting_id uuid.UUID
	last_id uuid.UUID

	start_event_key []byte
	end_event_key []byte

	stream_count int16

	chunk_size int16
	limit int16

	db *bolt.DB
}

func (self *BoltStream_) Reset() {

	self.last_id = self.starting_id
	self.stream_count = 0

	self.db.View(func(tx *bolt.Tx) error {

		events_log_index := tx.Bucket([]byte("events_index"))

		self.start_event_key = events_log_index.Get(self.starting_id.Bytes())

		if(self.start_event_key != nil) {

		}

		return nil
	})
}

func (self *BoltStream_) LastId() Identifier {
	return self.last_id
}

func (self *BoltStream_) Next() []Event {

	events := []*sourced.Event_{}

	start := binary.BigEndian.Uint64(self.start_event_key)
	end := start + self.chunk_size

	binary.BigEndian.PutUint64(self.end_event_key, end)

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

func (self *BoltStream_) Close() {
	self.db.Close()
}

func BoltStream(path string, starting_id uuid.UUID, chunk_size int16, limit int16) *BoltStream_ {

	db, err := bolt.Open(path, 0666, &bolt.Options{ReadOnly: true})
	if err != nil {
		log.Fatal(err)
	}

	bolt_stream := &BoltStream_ {
		starting_id: starting_id,
		last_id: starting_id,

		chunk_size: chunk_size,
		limit: limit,

		db: db,
	}

	bolt_stream.Reset()

	return bolt_stream
}
*/