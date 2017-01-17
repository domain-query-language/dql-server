package projector

import (
	"errors"
	"github.com/domain-query-language/dql-server/src/server/vm"
)

var (

	PROJECTOR_HANDLER_NOT_EXISTS = errors.New("Projector handler does not exist.")
)

type ProjectorHandler func(projection Projection, event vm.Event) Projection

type Projector interface {

	Reset()

	Apply(event vm.Event) error

	Projection() Projection

	Version() int

}

type SimpleProjector struct {

	version int

	projection Projection

	handlers map[vm.Identifier]ProjectorHandler

}

func (self *SimpleProjector) Reset() {
	self.version = 0
	self.projection.Reset()
}

func (self *SimpleProjector) Apply(event vm.Event) error {

	handler, ok := self.handlers[event.TypeId()]

	if(!ok) {
		return PROJECTOR_HANDLER_NOT_EXISTS
	}

	self.projection = handler(self.projection, event)
	self.version++

	return nil
}

func (self *SimpleProjector) State() Projection {
	return self.projection
}

func (self *SimpleProjector) Version() int {
	return self.version
}
