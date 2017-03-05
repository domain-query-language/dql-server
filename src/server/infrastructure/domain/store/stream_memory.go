package store

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/satori/go.uuid"
)

type MemoryStream struct {

	log *MemoryLog

	current_event vm.Loggable
	current int
}

func (self *MemoryStream) Reset() {

	self.current_event = nil
	self.current = -1
}

func (self *MemoryStream) LastId() vm.Identifier {

	if self.current_event == nil {
		return uuid.Nil
	}

	return self.current_event.Id()
}

func (self *MemoryStream) Seek(identifier vm.Identifier) {

	if identifier == uuid.Nil {

		self.Reset()

		return
	}

	index, exists := self.log.events_index[identifier]

	if exists {
		self.current = index
		self.current_event = self.log.events[index]

		return
	}

	self.current = len(self.log.events) - 1
	self.current_event = nil
}

func (self *MemoryStream) Next() bool {

	if(self.current >= (len(self.log.events) - 1)) {
		return false
	}

	self.current++
	self.current_event = self.log.events[self.current]

	return true
}

func (self *MemoryStream) Value() vm.Loggable {
	return self.current_event
}

func NewMemoryStream(log *MemoryLog) *MemoryStream {

	stream := &MemoryStream {
		log: log,
	}

	stream.Reset()

	return stream
}

type MemoryAggregateStream struct {

	id vm.AggregateIdentifier
	log *MemoryLog

	current_event vm.Loggable
	current int
}

func (self *MemoryAggregateStream) Reset() {

	self.current_event = nil
	self.current = -1
}

func (self *MemoryAggregateStream) Version() int {
	return self.current + 1
}

func (self *MemoryAggregateStream) Seek(version int) {

}

func (self *MemoryAggregateStream) Next() bool {

	if(self.current >= (len(self.log.aggregates[self.id]) - 1)) {
		return false
	}

	self.current++
	self.current_event = self.log.aggregates[self.id][self.current]

	return true
}

func (self *MemoryAggregateStream) Value() vm.Loggable {
	return self.current_event
}

func NewMemoryAggregateStream(id *vm.AggregateIdentifier, log *MemoryLog) *MemoryAggregateStream {

	stream := &MemoryAggregateStream {
		id: (*id),
		log: log,
	}

	stream.Reset()

	return stream
}
