package adapter

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

type Adapter interface {

	Next() (*Handleable, error)
}

// Make it easy to get the correct type out, no need for casting based on "HandleableType"
type Handleable struct {

	Typ HandleableType
	Command vm.Command
	Query vm.Query
}

type HandleableType string

const (
	CMD HandleableType = "command"
	QRY = "query"
)

// Helper methods to make creating it easier
func NewCommand(cmd vm.Command) *Handleable {

	return &Handleable{CMD, cmd, nil}
}

func NewQuery(qry vm.Query) *Handleable {

	return &Handleable{QRY, nil, qry}
}

func (h *Handleable) String() string {
	return string(h.Typ)
}
