package valueobjects

import (
	"regexp"
	"errors"
)

type Name string

var isName = regexp.MustCompile(`^[a-zA-Z\d._-]+$`).MatchString

func (n Name) Check() error {

	if (!isName(string(n))) {
		return errors.New("Invalid name: "+string(n))
	}

	if (len(string(n)) > 120) {
		return errors.New("Name is too long, max length is 120 characters")
	}
	return nil
}

func NewName(name string) (Name, error) {

	var vo Name = Name(name)
	return vo, vo.Check()
}
