package store

type Stream interface {

	Reset()

	LastId() Identifier

	Next() Event

}
