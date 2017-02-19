package aggregate

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/aggregate"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"errors"
	"github.com/domain-query-language/dql-server/src/server/domain/store"
)

type MemoryRepository struct {

	event_log store.Log
	command_log store.Log

	aggregates_archetypes map[vm.Identifier]aggregate.Aggregate
	aggregate_instances map[vm.AggregateIdentifier]aggregate.Aggregate
}

func (self *MemoryRepository) Add(aggregate aggregate.Aggregate) {

	aggregate.Reset()

	self.aggregates_archetypes[aggregate.Id().TypeId] = aggregate
}

func (self *MemoryRepository) Get(id *vm.AggregateIdentifier) (aggregate.Aggregate, error) {

	aggregate, instance_exists := self.aggregate_instances[(*id)]

	if instance_exists {
		return aggregate, nil
	}

	aggregate_archetype, found := self.aggregates_archetypes[id.TypeId]

	if !found {
		return nil, errors.New("The aggregate type does not exist.")
	}

	aggregate = aggregate_archetype.Copy(id.Id)

	stream := self.event_log.AggregateStream(id)

	for stream.Next() {
		aggregate.Apply(stream.Value().(vm.Event))
	}

	aggregate.Flush()

	return aggregate, nil
}

func (self *MemoryRepository) Save(aggregate aggregate.Aggregate) error {

	for _, event := range aggregate.Events() {
		self.event_log.Append(event)
	}

	for _, command := range aggregate.Commands() {
		self.command_log.Append(command)
	}

	aggregate.Flush()

	self.aggregate_instances[(*aggregate.Id())] = aggregate

	return nil
}

func CreateMemoryRepository(event_log store.Log, command_log store.Log) *MemoryRepository {

	return &MemoryRepository {
		event_log: event_log,
		command_log: command_log,
		aggregates_archetypes: map[vm.Identifier]aggregate.Aggregate{},
		aggregate_instances: map[vm.AggregateIdentifier]aggregate.Aggregate{},
	}
}
