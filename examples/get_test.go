package examples

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/wgarcia4190/go-rest/gorest_mock"
)

func TestMain(m *testing.M) {
	fmt.Println("About to start test cases for package 'examples'")

	// Tell the HTTP library to mock any further requests from here.
	gorest_mock.MockupServer.Start()

	os.Exit(m.Run())
}

func TestGetEndpoints(t *testing.T) {

	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		// Initialization
		gorest_mock.MockupServer.DeleteMocks()
		gorest_mock.MockupServer.AddMock(gorest_mock.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting github endpoints"),
		})

		// Execution
		endpoints, err := GetEndpoints()

		// Validation
		if endpoints != nil {
			t.Error("no endpoints expected at this point")
		}

		if err == nil {
			t.Error("ar error was expected")
		}

		if err != nil && err.Error() != "timeout getting github endpoints" {
			t.Error("invalid error message received")
		}

	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		// Initialization
		gorest_mock.MockupServer.DeleteMocks()
		gorest_mock.MockupServer.AddMock(gorest_mock.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url":123}`,
		})

		// Execution
		endpoints, err := GetEndpoints()

		// Validation
		if endpoints != nil {
			t.Error("no endpoints expected at this point")
		}

		if err == nil {
			t.Error("ar error was expected")
		}

		if err != nil && !strings.Contains(err.Error(), "cannot unmarshal") {
			t.Error("invalid error message received")
		}
	})

	t.Run("TestNoError", func(t *testing.T) {
		// Initialization
		gorest_mock.MockupServer.DeleteMocks()
		gorest_mock.MockupServer.AddMock(gorest_mock.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": "https://api.github.com/user"}`,
		})

		// Execution
		endpoints, err := GetEndpoints()

		// Validation
		if err != nil {
			t.Error(fmt.Sprintf("no error was excepted and we got %s", err.Error()))
		}

		if endpoints == nil {
			t.Error("endpoints were expected and we got nil")
		}

		if endpoints != nil && endpoints.CurrentUserUrl != "https://api.github.com/user" {
			t.Error("invalid current user url")
		}

	})
}
