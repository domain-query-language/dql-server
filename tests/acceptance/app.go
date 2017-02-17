package acceptance

import (
	"errors"
	"fmt"
	"bytes"
	"net/http/httptest"
	"encoding/json"
	"strings"
	controllers "github.com/domain-query-language/dql-server/examples/dql/application/http"
	"net/http"
	"github.com/domain-query-language/dql-server/examples/dql/infrastructure"
)

type app struct {

	server http.Handler

}

func NewApp() *app {

	infrastructure.Boot()
	handler := controllers.SetupServer()
	return &app{handler}
}


func (a *app) given(statements ...string) error {

	for _, statement := range statements {

		_, err := a.process(statement)

		if (err != nil) {
			return err
		}
	}

	return nil
}

func (a *app) process(statement string) (string, error) {

	response, err := a.makeHttpRequest(statement, http.StatusOK)

	if err != nil {
		return "", err
	}

	if (!strings.Contains(response, "data")) {
		msg := "Body does not data"
		return response, errors.New(msg)
	}

	return response, nil
}


func (a *app) processAndFail(statement string) (string, error) {

	response, err := a.makeHttpRequest(statement, http.StatusBadRequest)

	if err != nil {
		return "", err
	}

	if (!strings.Contains(response, "error")) {
		msg := "Body does not contain error"
		return response, errors.New(msg)
	}

	return response, nil
}

func (a *app) makeHttpRequest(statement string, status int) (string, error) {

	request, err := newRequest(statement)

	if err != nil {
		return "", err
	}

	response, st := a.handleRequest(request)

	if st != status {
		msg := fmt.Sprintf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		return "", errors.New(msg)
	}

	if (!isJSON(response)) {
		msg := "Response is not JSON"
		return response, errors.New(msg)
	}

	return response, nil
}

func newRequest(statement string) (*http.Request, error) {

	input := bytes.NewBuffer([]byte(statement))

	return http.NewRequest("POST", "/schema", input)
}

func (a *app) handleRequest(request *http.Request) (response string, code int) {

	responseRecorder := httptest.NewRecorder()

	a.server.ServeHTTP(responseRecorder, request)

	response = responseRecorder.Body.String()
	code = responseRecorder.Code

	return
}

func isJSON(s string) bool {

	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
