package vm

type AggregateIdentifier struct {

	Id Identifier `json:"id"`
	TypeId Identifier `json:"type_id"`

}

func (self *AggregateIdentifier) Bytes() []byte {
	return append(self.Id.Bytes(), self.TypeId.Bytes()...)
}

func NewAggregateIdentifier(id Identifier, type_id Identifier) *AggregateIdentifier {

	return &AggregateIdentifier {
		Id: id,
		TypeId: type_id,
	}
}
