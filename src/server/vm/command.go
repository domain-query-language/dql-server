package vm

type Command interface {

	Id() Identifier

	TypeId() Identifier

	AggregateTypeId() Identifier
}
