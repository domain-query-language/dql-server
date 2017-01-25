package vm

type Event interface {

	Id() Identifier

	TypeId() Identifier

	AggregateTypeId() Identifier
}
