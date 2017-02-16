package database

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

type Projection struct {

	Identifier vm.Identifier

	IsCreated bool
	IsDeleted bool
}

func (self Projection) Id() vm.Identifier {
	return self.Identifier
}

func (self Projection) Reset() {
	self.IsCreated = false
}

func (self *Projection) Create() {
	self.IsCreated = true
}

func (self *Projection) Delete() {
	self.IsDeleted = true
}

var State = Projection {
	Identifier: Identifier,
	IsCreated: false,
	IsDeleted: false,
}