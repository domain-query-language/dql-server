package store
/*
import (
	"github.com/boltdb/bolt"
	"encoding/binary"
	"log"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/src/server/domain/store"
)

type BoltLog_ struct {

	db *bolt.DB
}

func (self *BoltLog_) Reset() {

	self.db.Update(func(tx *bolt.Tx) error {

		tx.DeleteBucket([]byte("aggregates"))
		tx.CreateBucket([]byte("aggregates"))

		tx.DeleteBucket([]byte("events_index"))
		tx.CreateBucket([]byte("events_index"))

		tx.DeleteBucket([]byte("events_log"))
		tx.CreateBucket([]byte("events_log"))

		return nil
	})
}

func (self *BoltLog_) Append(events []vm.Loggable) {

	self.db.Update(func(tx *bolt.Tx) error {

		aggregates_bucket := tx.Bucket([]byte("aggregates"))
		aggregate_key := make([]byte, 8)
		aggregate_id := make([]byte, 32)

		events_log_bucket := tx.Bucket([]byte("events_log"))
		events_log_index := tx.Bucket([]byte("events_index"))
		event_key := make([]byte, 8)

		key_id := uint64(0)

		for _, event := range events {

			// Generating the Event key.
			key_id, _ = events_log_bucket.NextSequence()
			binary.BigEndian.PutUint64(event_key, key_id)

			// Adding an Event to the Log.
			events_log_bucket.Put(
				event_key,
				event.GobEncode(),
			)

			// Indexing the Event identifier.
			events_log_index.Put(
				event.Id().Bytes(),
				event_key,
			)

			//
			//	Associate with Aggregate
			//

			aggregate_id = append(event.AggregateId().Bytes(), event.AggregateId().Bytes()...)

			// Create a new Aggregate bucket.
			aggregate_bucket, _ := aggregates_bucket.CreateBucketIfNotExists(
				aggregate_id,
			)

			// Generating the next Aggregate Event key.
			key_id, _ = aggregate_bucket.NextSequence()
			binary.BigEndian.PutUint64(aggregate_key, key_id)

			// Adding an Event key reference to the Aggregate.
			aggregate_bucket.Put(
				aggregate_key,
				event_key,
			)
		}

		return nil
	})
}

func (self *BoltLog_) Stream() store.Stream {

}

func (self *BoltLog_) AggregateStream(id *vm.AggregateIdentifier) store.AggregateStream {

}

func (self *BoltLog_) Close() {
	self.db.Close()
}

func BoltLog(path string) *BoltLog_ {

	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &BoltLog_{
		db: db,
	}
}
*/