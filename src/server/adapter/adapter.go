package adapter

import (
	"github.com/domain-query-language/dql-server/src/server/vm/handler"
	"fmt"
)

type Adapter interface {

	Next() (*Handleable, error)
}

// Make it easy to get the correct type out, no need for casting based on "HandleableType"
type Handleable struct {

	Typ HandleableType
	Command handler.Command
	Query handler.Query
}

func (h *Handleable) String() string {
	typ := fmt.Sprintf("%v", h.Typ)
	return typ+": "+h.Query.String();
}

type HandleableType string

const (
	CMD HandleableType = "command"
	QRY = "query"
)

// Helper methods to make creating it easier
func NewCommand(cmd handler.Command) *Handleable {

	return &Handleable{CMD, cmd, nil}
}

func NewQuery(qry handler.Query) *Handleable {

	return &Handleable{QRY, nil, qry}
}

