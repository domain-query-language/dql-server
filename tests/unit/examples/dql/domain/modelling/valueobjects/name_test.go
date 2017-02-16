package valueobjects

import (
	"testing"
	"github.com/domain-query-language/dql-server/examples/dql/domain/modelling/valueobjects"
)

var valid = []string {
	"db",
	"db1",
	"db-1",
	"db_1",
	"db.1",
}

func TestValidNames(t *testing.T) {

	for _, value := range valid {

		name, err := valueobjects.NewName(value)

		if (err != nil) {
			t.Error("Value should be considered valid")
			t.Error(err.Error())
		}

		if string(name) != value {
			t.Error("Values do not match")
			t.Error("Expected: "+value)
			t.Error("Got: "+string(name))
		}
	}
}

var invalid = []string {
	"",
	"%^&#",
	"[db",
	"db]",
	"db@1",
}

func TestInvalidValidNames(t *testing.T) {

	for _, value := range invalid {

		_, err := valueobjects.NewName(value)

		if (err == nil) {
			t.Error("Invalid name acceptance")
			t.Error("value: "+value)
		}
	}
}
