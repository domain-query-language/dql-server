package vm

type Command interface {

	Id() Identifier

	TypeId() Identifier

	AggregateId() Identifier

	AggregateTypeId() Identifier
}
