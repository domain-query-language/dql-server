package acceptance

import (
	"testing"
	"net/http/httptest"
	"net/http"
	controllers "github.com/domain-query-language/dql-server/examples/dql/application/http"
	"bytes"
	"encoding/json"
)

const (
	LIST_DATABASE string = "LIST DATABASES;"
	CREATE_DATABASE = "CREATE DATABASE 'db1';"
)

func TestListDatabases(t *testing.T) {

	responseRecorder := httptest.NewRecorder()

	input := bytes.NewBuffer([]byte(LIST_DATABASE))

	req, err := http.NewRequest("POST", "/schema", input)

	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(controllers.Schema)

	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	response := responseRecorder.Body.String()

	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}


	if (!isJSON(response)) {
		t.Error("Response is not JSON")
		t.Error(response)
		return
	}

	//if (strings.Contains(response, ""))
}

func isJSON(s string) bool {

	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}