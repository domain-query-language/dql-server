package query

type Handler interface {

	Handle(query Query) (result Result, error)
}

