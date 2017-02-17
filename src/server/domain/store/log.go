package store

import "github.com/domain-query-language/dql-server/src/server/domain/vm"

type Log interface {

	Reset()

	Append(events []vm.Event)

	AppendCommands(commands []vm.Command)

	Stream() Stream

	AggregateStream(id *vm.AggregateIdentifier) AggregateStream
}
