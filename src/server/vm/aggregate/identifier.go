package aggregate

import "github.com/satori/go.uuid"

type Identifier interface {

	Bytes() []byte

}

type AggregateIdentifier struct {

	id uuid.UUID
	type_id uuid.UUID

}

func (self *AggregateIdentifier) Bytes() []byte {
	return append(self.type_id.Bytes(), self.id.Bytes()...)
}
