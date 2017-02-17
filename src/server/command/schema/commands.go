package schema

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

type CommandID string;

func (c CommandID) Bytes() []byte {

	return []byte(c);
}

/**
 * Command IDs
 */
const (
	CREATE_DATABASE CommandID = "47d6049c-251e-48ce-b4d4-38089208f556"
)


/**
 * Commands
 */
type CreateDatabase struct {
	Name string
}

func (c *CreateDatabase) String () string {

	return "CreateDatabase";
}

func (c *CreateDatabase) Id () vm.Identifier {

	return CREATE_DATABASE;
}

func (c *CreateDatabase) TypeId() vm.Identifier {

	return CREATE_DATABASE;
}

func (c *CreateDatabase) AggregateId() vm.Identifier {

	return CREATE_DATABASE;
}

func (c *CreateDatabase) ContextId() vm.Identifier {

	return CREATE_DATABASE;
}