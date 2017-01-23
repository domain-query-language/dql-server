package adapter

type Adapter interface {
	Next() (Handleable, error)
}

type Handleable interface {
	Type() (HandleableType)
}

type HandleableType string

const (
	CMD HandleableType = "command"
	EVT  = "event"
)

