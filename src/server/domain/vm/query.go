package vm

type Query interface {

	Id() Identifier
}

type Query_ struct {

	id Identifier
	Payload Payload
}

func (self *Query_) Id() Identifier {
	return self.id
}

func NewQuery(id Identifier, payload Payload) *Query_ {
	return &Query_{
		id: id,
		Payload: payload,
	}
}
