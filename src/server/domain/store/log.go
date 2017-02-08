package store

type Log interface {

	Reset()

	Append(events []Event)
}
