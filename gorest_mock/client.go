package gorest_mock

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type httpClientMock struct {
}

func (c *httpClientMock) Do(r *http.Request) (*http.Response, error) {
	requestBody, err := r.GetBody()

	if err != nil {
		return nil, err
	}
	defer func() {
		if err := requestBody.Close(); err != nil {
			fmt.Println("Error closing request")
		}
	}()

	body, err := ioutil.ReadAll(requestBody)

	if err != nil {
		return nil, err
	}

	var response http.Response

	mock := MockupServer.mocks[MockupServer.getMockKey(r.Method, r.URL.String(), string(body))]
	if mock != nil {
		if mock.Error != nil {
			return nil, mock.Error
		}
		response.StatusCode = mock.ResponseStatusCode
		response.Body = ioutil.NopCloser(strings.NewReader(mock.ResponseBody))
		response.ContentLength = int64(len(mock.ResponseBody))
		response.Request = r
		return &response, nil
	}
	return nil, fmt.Errorf("no mock matching %s from '%s' with given body", r.Method, r.URL.String())
}
