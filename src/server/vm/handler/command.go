package handler

type Command interface {

	Id() Identifier

	TypeId() Identifier
}
