package acceptance

import (
	"testing"
	"strings"
)


const (
	LIST_DATABASE string = "LIST DATABASES;"
	CREATE_DATABASE = "CREATE DATABASE 'db1';"
)

func TestListDatabases(t *testing.T) {

	app := NewApp()

	response, err := app.when(LIST_DATABASE)

	if (err != nil) {
		t.Error(err)
		return
	}

	expected := "data"

	if (!strings.Contains(response, expected)) {
		t.Error("Does not contain expected value '"+expected+"'")
		t.Error(response)
	}
}

func TestAddingDatabase(t *testing.T) {

	app := NewApp()

	err := app.given(CREATE_DATABASE)

	if (err != nil) {
		t.Error(err)
		return
	}

	response, err := app.when(LIST_DATABASE)

	if (err != nil) {
		t.Error(err)
		return
	}

	expected := "db1"

	if (!strings.Contains(response, expected)) {
		t.Error("Does not contain expected value '"+expected+"'")
		t.Error(response)
	}
}

func TestCannotAddDatabaseThatAlreadyExists(t *testing.T) {

	app := NewApp()

	err := app.given(CREATE_DATABASE)

	if (err != nil) {
		t.Error(err)
		return
	}

	_, err = app.whenError(CREATE_DATABASE)

	if (err != nil) {
		t.Error(err)
		return
	}
}

func TestInvalidStatement(t *testing.T) {

	app := NewApp()

	_, err := app.whenError("CREATE DATABASE '[[[--%%%';")

	if (err != nil) {
		t.Error(err)
		return
	}
}