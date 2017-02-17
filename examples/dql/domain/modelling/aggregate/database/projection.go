package database

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
)

type Projection struct {

	Identifier vm.Identifier

	IsCreated bool
	IsDeleted bool
}

func (self *Projection) Id() vm.Identifier {
	return self.Identifier
}

func (self *Projection) Reset() {
	self.IsCreated = false
	self.IsDeleted = false
}

func (self *Projection) Copy() projection.Projection {

	projection := *self

	return &projection
}

func (self *Projection) Create() {
	self.IsCreated = true
}

func (self *Projection) Delete() {
	self.IsDeleted = true
}

var State = &Projection {
	Identifier: Identifier,
	IsCreated: false,
	IsDeleted: false,
}
