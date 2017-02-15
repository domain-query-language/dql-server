package list_databases

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"sort"
	"github.com/domain-query-language/dql-server/examples/dql/application/projection/list-databases"
	"github.com/satori/go.uuid"
)

type ListDatabases struct {

	Pid vm.Identifier
	Databases map[vm.Identifier]string
}

func (self *ListDatabases) Id() vm.Identifier {
	return self.Pid
}

func (self *ListDatabases) Reset() {
	self.Databases = map[vm.Identifier]string{}
}

func (self *ListDatabases) Add(id vm.Identifier, name string) {
	self.Databases[id] = name
}

func (self *ListDatabases) Rename(id vm.Identifier, name string) {
	self.Databases[id] = name
}

func (self *ListDatabases) Remove(id vm.Identifier) {
	delete(self.Databases, id)
}

func (self *ListDatabases) List() []string {

	names := []string{}

	for _, v := range self.Databases {
		names = append(names, v)
	}

	sort.Strings(names)

	return names
}

var Projection = &ListDatabases{
	Pid: list_databases.Identifier,
	Databases: map[vm.Identifier]string {
		uuid.NewV4(): "master-0.0.1",
		uuid.NewV4(): "master-0.0.3",
		uuid.NewV4(): "master-0.0.2",
		uuid.NewV4(): "schema",
	},
}
