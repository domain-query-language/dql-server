package player

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"time"
	"github.com/satori/go.uuid"
)

type Snapshot struct {

	Id vm.Identifier

	ContextId vm.Identifier

	LastId vm.Identifier

	OccurredAt time.Time

	Version int
}

func NewSnapshot(identifier vm.Identifier, contextId vm.Identifier) *Snapshot {

	return &Snapshot {
		Id: identifier,
		ContextId: contextId,
		LastId: uuid.Nil,
		OccurredAt: time.Now(),
		Version: 0,
	}
}
