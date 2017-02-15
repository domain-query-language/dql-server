package store

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/satori/go.uuid"
	"testing"
	"github.com/domain-query-language/dql-server/src/server/infrastructure/domain/store"
	store_contract "github.com/domain-query-language/dql-server/src/server/domain/store"
)

var Payload = struct {
	Name string
	Value int
}{
	Name: "Colin",
	Value: 1,
}

var AGGREGATE_A = uuid.NewV4()
var AGGREGATE_B = uuid.NewV4()

var AGGREGATE_TYPE_A = uuid.NewV4()
var AGGREGATE_TYPE_B = uuid.NewV4()

var events = []vm.Event {
	vm.NewEvent(
		uuid.FromStringOrNil("77a15dde-ef61-4699-a609-f5b0f2e2cb3e"),
		uuid.FromStringOrNil("2ad1fe94-b3d4-4fd5-abfc-0b1b09564f1b"),
		AGGREGATE_A,
		AGGREGATE_TYPE_A,
		Payload,
	),
	vm.NewEvent(
		uuid.FromStringOrNil("7897a07d-c28f-4292-b712-82da544723d1"),
		uuid.FromStringOrNil("3b4d9994-5603-4082-a411-74b7be29ff09"),
		AGGREGATE_A,
		AGGREGATE_TYPE_A,
		Payload,
	),
	vm.NewEvent(
		uuid.FromStringOrNil("73dc3f5e-adeb-419f-a9e6-887c56e1535d"),
		uuid.FromStringOrNil("91009b64-ae12-43f0-8cf3-9d4cbbc676b9"),
		AGGREGATE_A,
		AGGREGATE_TYPE_A,
		Payload,
	),
	vm.NewEvent(
		uuid.FromStringOrNil("c5ed02e2-646c-4687-b7d3-3ec68b81a72e"),
		uuid.FromStringOrNil("5ceb80bd-e48c-4a9f-b244-a3cda4037ab9"),
		AGGREGATE_B,
		AGGREGATE_TYPE_A,
		Payload,
	),
	vm.NewEvent(
		uuid.FromStringOrNil("9ffe4ffc-b2ba-46c5-9326-bc66294e9e15"),
		uuid.FromStringOrNil("0bc8f850-c8b5-4375-a751-b36560ce8422"),
		AGGREGATE_B,
		AGGREGATE_TYPE_A,
		Payload,
	),
}

var eventTests = []struct {
	log store_contract.Log
	events []vm.Event
	expected_identifiers []vm.Identifier
}{
	{
		store.NewMemoryLog(),
		[]vm.Event{},
		[]vm.Identifier{},
	},
	{
		store.NewMemoryLog(),
		events,
		[]vm.Identifier{
			events[0].Id(),
			events[1].Id(),
			events[2].Id(),
			events[3].Id(),
			events[4].Id(),
		},
	},
}

func TestMemoryStore(t *testing.T) {

	for _, tt := range eventTests {

		tt.log.Append(tt.events)

		stream := tt.log.Stream()

		if(stream.LastId() != uuid.Nil) {
			t.Errorf(
				"Stream.LastId(): expected %v, got %v.",
				uuid.Nil,
				stream.LastId(),
			)
		}

		for i :=0; stream.Next(); i++ {

			event := stream.Value()

			if(stream.LastId() != event.Id()) {
				t.Errorf(
					"Stream.LastId(): expected %v, got %v.",
					event.Id(),
					stream.LastId(),
				)
			}

			if(event.Id() != tt.expected_identifiers[i]) {
				t.Errorf(
					"Stream.Value(): expected %v, got %v.",
					tt.expected_identifiers[i],
					event.Id(),
				)
			}
		}
	}
}

func TestMemoryStoreAfterLogReset(t *testing.T) {

	for _, tt := range eventTests {

		tt.log.Append(tt.events)
		tt.log.Reset()
		tt.log.Append(tt.events)

		stream := tt.log.Stream()

		if(stream.LastId() != uuid.Nil) {
			t.Errorf(
				"Stream.LastId(): expected %v, got %v.",
				uuid.Nil,
				stream.LastId(),
			)
		}

		for i :=0; stream.Next(); i++ {

			event := stream.Value()

			if(stream.LastId() != event.Id()) {
				t.Errorf(
					"Stream.LastId(): expected %v, got %v.",
					event.Id(),
					stream.LastId(),
				)
			}

			if(event.Id() != tt.expected_identifiers[i]) {
				t.Errorf(
					"Stream.Value(): expected %v, got %v.",
					tt.expected_identifiers[i],
					event.Id(),
				)
			}
		}
	}
}

func TestMemoryStoreAfterStreamReset(t *testing.T) {

	for _, tt := range eventTests {

		tt.log.Append(tt.events)

		stream := tt.log.Stream()

		if(stream.LastId() != uuid.Nil) {
			t.Errorf(
				"Stream.LastId(): expected %v, got %v.",
				uuid.Nil,
				stream.LastId(),
			)
		}

		for i :=0; stream.Next(); i++ {}

		stream.Reset()

		for i :=0; stream.Next(); i++ {

			event := stream.Value()

			if(stream.LastId() != event.Id()) {
				t.Errorf(
					"Stream.LastId(): expected %v, got %v.",
					event.Id(),
					stream.LastId(),
				)
			}

			if(event.Id() != tt.expected_identifiers[i]) {
				t.Errorf(
					"Stream.Value(): expected %v, got %v.",
					tt.expected_identifiers[i],
					event.Id(),
				)
			}
		}
	}
}
