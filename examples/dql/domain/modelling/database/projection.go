package database

import "github.com/domain-query-language/dql-server/src/server/domain/vm"

type Projection struct {

	Id vm.Identifier

	IsCreated bool
}

func (self *Projection) Id() vm.Identifier {
	return self.Id
}

func (self *Projection) Reset() {
	self.IsCreated = false
}

func (self *Projection) Create() {
	self.IsCreated = true
}

var Projection = Projection {
	Id: Identifier,
	IsCreated: false,
}
