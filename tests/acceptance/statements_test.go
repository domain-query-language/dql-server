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
)

const (
	LIST_DATABASE string = "LIST DATABASES;"
	CREATE_DATABASE = "CREATE DATABASE 'db1';"
)

func TestListDatabases(t *testing.T) {

	infrastructure.Boot()

	responseRecorder := httptest.NewRecorder()

	input := bytes.NewBuffer([]byte(LIST_DATABASE))

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
		return
	}

	//if (strings.Contains(response, ""))
}

func isJSON(s string) bool {

	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}