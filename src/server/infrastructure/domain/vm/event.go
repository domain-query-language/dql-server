package vm

import "github.com/domain-query-language/dql-server/src/server/domain/vm"

type Event struct {

	id vm.Identifier

	type_id vm.Identifier

	context_id vm.Identifier

	payload struct{}
}

func NewEvent(id vm.Identifier, type_id vm.Identifier) *Event {

}
