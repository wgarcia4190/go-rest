package examples

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/wgarcia4190/go-rest/gorest_mock"
)

func TestCreateRepo(t *testing.T) {
	t.Run("TimeoutFromGithub", func(t *testing.T) {
		gorest_mock.MockupServer.DeleteMocks()
		gorest_mock.MockupServer.AddMock(gorest_mock.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: `{"name": "test-repo", "private": true}`,
			Error:       errors.New("timeout from github"),
		})

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}

		repo, err := CreateRepo(repository)

		if repo != nil {
			t.Error("no repo expected when we get a timeout from github")
		}

		if err == nil {
			t.Error("an error is expected when we get a timeout from github")
		}

		if err != nil && err.Error() != "timeout from github" {
			fmt.Println(err.Error())
			t.Error("invalid error message")
		}
	})

	t.Run("NoError", func(t *testing.T) {
		gorest_mock.MockupServer.DeleteMocks()
		gorest_mock.MockupServer.AddMock(gorest_mock.Mock{
			Method:             http.MethodPost,
			Url:                "https://api.github.com/user/repos",
			RequestBody:        `{"name": "test-repo", "private": true}`,
			ResponseStatusCode: http.StatusCreated,
			ResponseBody:       `{"id": 123, "name": "test-repo"}`,
		})

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}

		repo, err := CreateRepo(repository)

		if err != nil {
			fmt.Println(err.Error())
			t.Error("no error expected when we get a valid response")
		}

		if repo == nil {
			t.Error("a valid repo was expected")
		}

		if repo != nil && repo.Name != repository.Name {
			t.Error("invalid repository name obtained from github")
		}
	})
}
