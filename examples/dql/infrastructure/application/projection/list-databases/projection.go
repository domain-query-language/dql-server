package list_databases

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"sort"
	"github.com/domain-query-language/dql-server/examples/dql/application/projection/list-databases"
)

type ListDatabases struct {

	Pid vm.Identifier
	Databases []string
}

func (self *ListDatabases) Id() vm.Identifier {
	return self.Pid
}

func (self *ListDatabases) Reset() {
	self.Databases = []string{}
}

func (self *ListDatabases) Add(name string) {
	self.Databases = append(self.Databases, name)
}

func (self *ListDatabases) Remove(name string) {
	for i := range self.Databases {
		if self.Databases[i] == name {
			self.Databases = append(self.Databases[:i], self.Databases[i+1:]...)
		}
	}
}

func (self *ListDatabases) List() []string {

	sort.Strings(self.Databases)

	return self.Databases
}

var Projection = &ListDatabases{
	Pid: list_databases.Identifier,
	Databases: []string {
		"master-0.0.1",
		"master-0.0.3",
		"master-0.0.2",
		"schema",
	},
}
