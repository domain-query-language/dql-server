package database_name_unique

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/projection/database-name-unique"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
)

type DatabaseNameUnique struct {

	Pid vm.Identifier
	Databases map[vm.Identifier]string
}

func (self *DatabaseNameUnique) Id() vm.Identifier {
	return self.Pid
}

func (self *DatabaseNameUnique) Reset() {
	self.Databases = map[vm.Identifier]string{}
}

func (self *DatabaseNameUnique) Copy() projection.Projection {

	projection := *self

	return &projection
}

func (self *DatabaseNameUnique) Add(id vm.Identifier, name string) {
	self.Databases[id] = name
}

func (self *DatabaseNameUnique) Rename(id vm.Identifier, name string) {
	self.Databases[id] = name
}

func (self *DatabaseNameUnique) Remove(id vm.Identifier) {
	delete(self.Databases, id)
}

func (self *DatabaseNameUnique) Exists(name string) bool {

	for _, v := range self.Databases {
		if v == name {
			return true
		}
	}

	return false
}

var Projection = &DatabaseNameUnique {
	Pid: database_name_unique.Identifier,
	Databases: map[vm.Identifier]string {},
}
