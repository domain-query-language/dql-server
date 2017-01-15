package handler

type Handler interface {

	Handle(command Command) ([]Event, error)
}
