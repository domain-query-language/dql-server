package store

import "github.com/domain-query-language/dql-server/src/server/domain/vm"

type Log interface {

	Reset()

	Append(loggable vm.Loggable)

	Stream() Stream

	AggregateStream(id *vm.AggregateIdentifier) AggregateStream
}
