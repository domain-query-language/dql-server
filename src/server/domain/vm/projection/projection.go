package projection

import "github.com/domain-query-language/dql-server/src/server/domain/vm"

type Projection interface {

	Id() vm.Identifier

	Reset()

	/*
	GobEncode() []byte

	GobDecode([]byte) error
	*/
}

