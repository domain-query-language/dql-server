package acceptance

import (
	"testing"
	"net/http/httptest"
	"net/http"
	controllers "github.com/domain-query-language/dql-server/examples/dql/application/http"
	"bytes"
	"encoding/json"
	"github.com/domain-query-language/dql-server/examples/dql/infrastructure"
	"strings"
	"errors"
	"fmt"
)

func given(statements ...string) error {

	for _, statement := range statements {

		_, err := processStatement(statement)

		if (err != nil) {
			return err
		}
	}

	return nil
}

func when(statement string) (string, error) {

	return processStatement(statement)
}

func whenError(statement string, errorCode int) (string, error) {

	request, err := makeRequest(statement)

	if err != nil {
		return "", errors.New(err.Error())
	}

	response, status := handleRequest(request)

	if status != errorCode {
		msg := fmt.Sprintf("handler returned wrong status code: got %v want %v", status, errorCode)
		return "", errors.New(msg)
	}

	if (!isJSON(response)) {
		msg := "Response is not JSON"
		return response, errors.New(msg)
	}

	if (!strings.Contains(response, "error")) {
		msg := "Body does not contain error"
		return response, errors.New(msg)
	}

	return response, nil
}

func processStatement(statement string) (string, error) {

	request, err := makeRequest(statement)

	if err != nil {
		return "", errors.New(err.Error())
	}

	response, status := handleRequest(request)

	if status != http.StatusOK {
		msg := fmt.Sprintf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		return "", errors.New(msg)
	}

	if (!isJSON(response)) {
		msg := "Response is not JSON"
		return response, errors.New(msg)
	}

	return response, nil
}

func handleRequest(request *http.Request) (response string, code int) {

	responseRecorder := httptest.NewRecorder()

	server := controllers.SetupServer()

	server.ServeHTTP(responseRecorder, request)

	response = responseRecorder.Body.String()
	code = responseRecorder.Code

	return
}

func makeRequest(statement string) (*http.Request, error) {

	input := bytes.NewBuffer([]byte(statement))

	return http.NewRequest("POST", "/schema", input)
}

func isJSON(s string) bool {

	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

const (
	LIST_DATABASE string = "LIST DATABASES;"
	CREATE_DATABASE = "CREATE DATABASE 'db1';"
)

func TestListDatabases(t *testing.T) {

	infrastructure.Boot()

	response, err := when(LIST_DATABASE)

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

	infrastructure.Boot()

	err := given(CREATE_DATABASE)

	if (err != nil) {
		t.Error(err)
		return
	}

	response, err := when(LIST_DATABASE)

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

	infrastructure.Boot()

	err := given(CREATE_DATABASE)

	if (err != nil) {
		t.Error(err)
		return
	}

	_, err = whenError(CREATE_DATABASE, http.StatusBadRequest)

	if (err != nil) {
		t.Error(err)
		return
	}
}

func TestInvalidStatement(t *testing.T) {

	_, err := whenError("CREATE DATABASE '[[[--%%%';", http.StatusBadRequest)

	if (err != nil) {
		t.Error(err)
		return
	}
}