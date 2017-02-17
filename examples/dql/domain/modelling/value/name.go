package value

import "errors"

type Name struct {

	Value string

	errors []errors
}

func (self *Name) Check() bool {

	if len(self.Value) >= 256 {
		self.errors = append(self.errors, errors.New("The maxiumum length of a database name is 256 characters."))
	}

	if len(self.Value) == 0 {
		self.errors = append(self.errors, errors.New("The database name must not be blank."))
	}

	return len(self.errors) == 0
}

func (self *Name) Errors() []errors {
	return self.errors
}

func NewName(value string) *Name {

	return &Name {
		Value: value,
	}
}
