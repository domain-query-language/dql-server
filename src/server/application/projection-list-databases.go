package application

import (
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/handler/query"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"sort"
)

var ListDatabasesQueryID, _ = uuid.FromString("c50d6791-3fc5-4be8-91fc-c01f20526872")
var ListDatabasesQuery = query.CreateQuery(ListDatabasesQueryID)
var ListDatabasesQueryHandler = func(query query.Query, p projection.Projection) query.Result {

	result := query.Result_{
		Data: p.(Projection).Databases,
	}

	return result
}

var ListDatabasesProjection = Projection {
	Pid: ListDatabasesQueryID,
	Databases: []string{
		"master-0.0.1",
		"master-0.0.3",
		"master-0.0.2",
		"schema",
	},
}

type Projection struct {

	Pid vm.Identifier
	Databases []string
}

func (self *Projection) Id() vm.Identifier {
	return self.Pid
}

func (self *Projection) Reset() {
	self.Databases = []string{}
}

func (self *Projection) Add(database string) {

	self.Databases = append(self.Databases, database)

	sort.Strings(self.Databases)
}
