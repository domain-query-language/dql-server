package aggregate

type Aggregate interface {

	Reset()

	Flush()

	State() State

	Changes() []Event
}
