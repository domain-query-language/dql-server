package acceptance

import (
	"testing"
	"net/http/httptest"
	"net/http"
	controllers "github.com/domain-query-language/dql-server/examples/dql/application/http"
	"bytes"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/domain-query-language/dql-server/examples/dql/infrastructure"
	"strings"
)

const (
	LIST_DATABASE string = "LIST DATABASES;"
	CREATE_DATABASE = "CREATE DATABASE 'db1';"
)

func TestListDatabases(t *testing.T) {

	infrastructure.Boot()

	response := makeRequest(LIST_DATABASE, t)

	expected := "data"

	if (!strings.Contains(response, expected)) {
		t.Error("Does not contain expected value '"+expected+"'")
		t.Error(response)
	}
}

func makeRequest(statement string, t *testing.T) string {

	responseRecorder := httptest.NewRecorder()

	input := bytes.NewBuffer([]byte(statement))

	request, err := http.NewRequest("POST", "/", input)

	if err != nil {
		t.Fatal(err)
	}

	controllers.Schema(responseRecorder, request)

	status := responseRecorder.Code
	response := responseRecorder.Body.String()

	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)

		t.Error(spew.Sdump(responseRecorder))
	}

	if (!isJSON(response)) {
		t.Error("Response is not JSON")
		t.Error(response)
	}

	return response
}

func isJSON(s string) bool {

	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

func TestAddingDatabase(t *testing.T) {

	infrastructure.Boot()

	makeRequest(CREATE_DATABASE, t)

	response := makeRequest(LIST_DATABASE, t)

	expected := "db1"

	if (!strings.Contains(response, expected)) {
		t.Error("Does not contain expected value '"+expected+"'")
		t.Error(response)
	}
}