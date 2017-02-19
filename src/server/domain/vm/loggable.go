package vm

type Loggable interface {

	Id() Identifier

	AggregateId() *AggregateIdentifier
}
