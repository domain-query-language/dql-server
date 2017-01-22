package store

type Snapshot struct {

	id Identifier

	version int

	event Event
}
