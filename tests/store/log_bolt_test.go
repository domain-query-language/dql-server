package store

import (
	"testing"
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/src/server/domain/store"
	"github.com/domain-query-language/dql-server/sourced"
)

var aggregate_id uuid.UUID
var aggregate_type_id uuid.UUID

func generate_events(aggregates_count int, per_aggregate_count int) []store.Event {

	events := []store.Event{}

	for i := 0; i < aggregates_count; i++ {

		aggregate_id = uuid.NewV4()
		aggregate_type_id = uuid.NewV4()

		for j := 0; j < per_aggregate_count; j++ {

			events = append(
				events,
				sourced.Event(
					uuid.NewV4(),
					aggregate_id,
					aggregate_type_id,
					1,
					make([]byte, 1000),
				),
			)

		}
	}

	return events
}

func BenchmarkAppend(b *testing.B) {

	b.StopTimer()

	log := store.BoltLog("test.db")
	log.Reset()

	batches := [][]store.Event {
		generate_events(10, 1000),
	}

	b.StartTimer()

	for _, events := range batches {
		log.Append(events)
	}

	b.StopTimer()

	log.Close()
}
