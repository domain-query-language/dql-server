package main

import (
	"github.com/satori/go.uuid"
	"fmt"
	"github.com/domain-query-language/dql-server/sourced"
)

/**

	USE CASES

	#1 - Getting Events from the Global Stream

		- The developer should be able to define the starting uuid.
			- If UUID is nil, it should start from the beginning of the sourced.
		- The developer should be able to set a limit to the number of events returned.

	#2 - Getting Events from an Aggregate Stream

		- The development should be able to define an starting event number.
			- If the starting event number is nil, it should start from the first event belonging to the aggregate.
		- The developer should be able to set a limit to the number of events returned.

	#3 - Getting Events from a Custom Event Stream

 */

/*
func reset(db *bolt.DB) {

	db.Update(func(tx *bolt.Tx) error {

		// Bootstrap Buckets
		tx.DeleteBucket([]byte("aggregates"))
		tx.CreateBucket([]byte("aggregates"))

		tx.DeleteBucket([]byte("aggregates_index"))
		tx.CreateBucket([]byte("aggregates_index"))

		//tx.DeleteBucket([]byte("streams"))
		//tx.CreateBucket([]byte("streams"))

		tx.DeleteBucket([]byte("events_index"))
		tx.CreateBucket([]byte("events_index"))

		tx.DeleteBucket([]byte("events_log"))
		tx.CreateBucket([]byte("events_log"))

		// Create Global Stream
		//global_stream, _ := tx.Bucket([]byte("streams")).CreateBucketIfNotExists(uuid.Nil.Bytes())
		//global_index, _ := global_stream.CreateBucketIfNotExists("index")

		return nil
	})
}

func gen_next_key(bucket *bolt.Bucket) []byte {
	key := make([]byte, 8)

	id, _ := bucket.NextSequence()

	binary.BigEndian.PutUint64(
		key,
		id,
	)

	return key
}

func gen_aggregate_data(db *bolt.DB) {

	db.Update(func(tx *bolt.Tx) error {

		aggregates_index_bucket := tx.Bucket([]byte("aggregates"))
		aggregates_bucket := tx.Bucket([]byte("aggregates"))
		events_index_bucket := tx.Bucket([]byte("events_index"))
		events_log_bucket := tx.Bucket([]byte("events_log"))

		aggregate_id := uuid.NewV4().Bytes()

		for i := 0; i < 100000; i++ {

			event_key := gen_next_key(events_log_bucket)

			events_log_bucket.Put(
				event_key,
				[]byte(strconv.Itoa(i)),
			)

			events_index_bucket.Put(
				uuid.NewV4().Bytes(),
				event_key,
			)

			aggregates_key := gen_next_key(aggregates_bucket)

			aggregate_bucket, _ := aggregates_bucket.CreateBucketIfNotExists(aggregates_key)
			aggregate_key := gen_next_key(aggregate_bucket)

			aggregate_bucket.Put(
				aggregate_key,
				event_key,
			)

			aggregates_index_bucket.Put(
				aggregate_id,
				aggregates_key,
			)
		}

		return nil
	})
}

func view_global_stream(db *bolt.DB, start_id uuid.UUID, limit int) {

	db.View(func(tx *bolt.Tx) error {

		events_index_bucket := tx.Bucket([]byte("events_index"))
		events_log_bucket := tx.Bucket([]byte("events_log"))
		events_log_cursor := events_log_bucket.Cursor()

		event_id := []byte(nil)

		if(!uuid.Equal(start_id, uuid.Nil)) {
			event_id = events_index_bucket.Get(start_id.Bytes())
		}

		i := 0

		fmt.Printf("Seeking to %v.\n", event_id)

		for k, v := events_log_cursor.Seek(event_id); k != nil; k, v = events_log_cursor.Next() {

			if(i >= limit) {
				return nil
			}

			key := binary.BigEndian.Uint64(k)
			val, _ := strconv.Atoi(string(v))

			fmt.Printf("key=%v, value=%v\n", key, val)

			i++
		}

		return nil
	})
}

func view_aggregate_stream(db *bolt.DB, aggregate_id uuid.UUID, offset uint64, limit int) {

	db.View(func(tx *bolt.Tx) error {

		aggregates_bucket := tx.Bucket([]byte("aggregates"))
		events_log_bucket := tx.Bucket([]byte("events_log"))

		aggregate_bucket := aggregates_bucket.Bucket(aggregate_id.Bytes())
		aggregate_bucket_cursor := aggregate_bucket.Cursor()

		fmt.Printf("Getting Aggregate %v.\n", aggregate_id)

		i := 0

		offset_key := make([]byte, 8)
		binary.BigEndian.PutUint64(
			offset_key,
			offset,
		)

		for k, v := aggregate_bucket_cursor.Seek(offset_key); k != nil; k, v = aggregate_bucket_cursor.Next() {

			if(i >= limit) {
				return nil
			}

			key := binary.BigEndian.Uint64(k)
			val := events_log_bucket.Get(v)

			fmt.Printf("key=%v, value=%v\n", key, val)

			i++
		}

		return nil
	})
}
*/

func main() {

	/*
	db, err := bolt.Open("events.db", 0600, nil)
	if err != nil {
		sourced.Fatal(err)
	}
	defer db.Close()

	reset(db)

	gen_aggregate_data(db)
	//gen_aggregate_data(db)
	//gen_aggregate_data(db)

	// Use Case #1

	//start_id := uuid.Nil
	//start_id, _ = uuid.FromString("0176b9ff-0c92-4055-b5bc-d3106b81f85b")
	//view_global_stream(db, start_id, 10)

	// Use Case #2

	//aggregate_id, _ := uuid.FromString("04e4aa0d-af9d-4bde-8a11-d973100d2d29")

	//view_aggregate_stream(db, aggregate_id, 0, 100)
	*/
}