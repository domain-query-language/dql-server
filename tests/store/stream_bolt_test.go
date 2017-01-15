package store

/*

import (
	"testing"
	"github.com/domain-query-language/dql-server/sourced"
	"github.com/domain-query-language/dql-server/stream"
	"github.com/satori/go.uuid"
	"fmt"
)

func BenchmarkStream(b *testing.B) {

	b.StopTimer()

	log := sourced.BoltLog("test.db")
	log.Reset()

	batches := [][]*sourced.Event_{
		generate_events(1, 1000),
		generate_events(1, 1000),
		generate_events(1, 1000),
		generate_events(1, 1000),
		generate_events(1, 1000),
		generate_events(1, 1000),
		generate_events(1, 1000),
		generate_events(1, 1000),
		generate_events(1, 1000),
		generate_events(1, 1000),

	}

	for _, events := range batches {
		log.Append(events)
	}

	log.Close()

	b.StartTimer()

	stream := stream.BoltStream("test.db", uuid.Nil, 1000, 10000)

	for events := stream.Next(); events != nil;  {
		for event := range events  {
			fmt.Printf("%v\n", event)
		}
	}


}
*/
