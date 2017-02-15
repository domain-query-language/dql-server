package projection

import (
	"errors"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

var (

	PROJECTOR_HANDLER_NOT_EXISTS = errors.New("Projector command does not exist.")
)

type ProjectorHandler func(projection Projection, event vm.Event) Projection

type Projector interface {

	Reset()

	Apply(event vm.Event) error

	Projection() Projection

	Version() int
}

/*
	Implementation of Projector
 */

type Projector_ struct {

	version int

	projection Projection

	handlers *map[vm.Identifier]ProjectorHandler
}

func (self *Projector_) Reset() {
	self.version = 0
	self.projection.Reset()
}

func (self *Projector_) Apply(event vm.Event) error {

	handler, ok := (*self.handlers)[event.TypeId()]

	if(!ok) {
		return PROJECTOR_HANDLER_NOT_EXISTS
	}

	self.projection = handler(self.projection, event)
	self.version++

	return nil
}

func (self *Projector_) Projection() Projection {
	return self.projection
}

func (self *Projector_) Version() int {
	return self.version
}

func NewProjector(projection Projection, handlers *map[vm.Identifier]ProjectorHandler) Projector_ {
	return Projector_ {
		version: 0,
		projection: projection,
		handlers: handlers,
	}
}
