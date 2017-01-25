package query

import "encoding/json"

type Result interface {

	MarshalJSON() ([]byte, error)
}

type Result_ struct {

	Data struct{}
}

func (self *Result_) MarshalJSON() ([]byte, error) {

	return json.Marshal(self.Data)
}
