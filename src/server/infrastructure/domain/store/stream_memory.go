package store

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/satori/go.uuid"
)

type MemoryStream struct {

	log *MemoryLog

	current_event vm.Event
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

func (self *MemoryStream) Value() vm.Event {
	return self.current_event
}

func NewMemoryStream(log *MemoryLog) *MemoryStream {

	stream := &MemoryStream {
		log: log,
	}

	stream.Reset()

	return stream
}
