package projector

import "github.com/domain-query-language/dql-server/src/server/vm"

type Projection interface {

	Reset()

	GobEncode() []byte

	GobDecode([]byte) error
}

type SimpleProjecton struct {

	values map[string]vm.Value
	entities map[vm.Identifier]vm.Entity
	indexes map[string]vm.Index
}
