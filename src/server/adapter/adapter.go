package adapter

import (
	"github.com/domain-query-language/dql-server/src/server/vm/handler"
)

type Adapter interface {

	Next() (Handleable, error)
}

// Make it easy to get the correct type out, no need for casting based on "HandleableType"
type Handleable interface {

	Type() HandleableType
	Command() handler.Command
	Query() handler.Query
}

type HandleableType string

const (
	CMD HandleableType = "command"
	QRY = "query"
)

// Simple implementation
type SimpleHandleable struct {

	typ HandleableType
	cmd handler.Command
	qry handler.Query
}

func (h *SimpleHandleable) Type() {

	return h.typ;
}

func (h *SimpleHandleable) Command() handler.Command {

	return h.cmd
}

func (h *SimpleHandleable) Query() handler.Query {

	return h.qry
}

// Helper methods to make creating it easier
func NewCommand(cmd handler.Command) Handleable {

	return &SimpleHandleable{CMD, cmd, nil}
}

func NewQuery(qry handler.Query) Handleable {

	return &SimpleHandleable{QRY, qry, nil}
}

