package aggregate

import "time"

type Snapshot struct {

	Id Identifier

	OccurredAt time.Time

	Version int

	Payload []byte
}

func (self *Snapshot) Encode() []byte {
	return []byte{}
}
