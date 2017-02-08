package query

import "github.com/domain-query-language/dql-server/src/server/domain/vm"

type Query interface {

	Id() vm.Identifier
}

type Query_ struct {
	id vm.Identifier
}

func (self *Query_) Id() vm.Identifier {
	return self.id
}

func CreateQuery(id vm.Identifier) Query {
	return &Query_{
		id: id,
	}
}
