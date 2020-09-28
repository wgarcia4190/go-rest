package gorest_mock

import (
	"fmt"
	"net/http"

	"github.com/wgarcia4190/go-rest/core"
)

// Mock structure provides a clean way to configure HTTP mocks based on
// the combination between request method, URL and request body.
type Mock struct {
	Method      string
	Url         string
	RequestBody string

	ResponseBody       string
	ResponseStatusCode int
	Error              error
	ResponseHeaders    http.Header
}

// GetResponse returns a Response object based on the mock configuration.
func (m *Mock) GetResponse() (*core.Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	// Fill response object with current mock details:
	response := core.Response{
		Status:     fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
		StatusCode: m.ResponseStatusCode,
		Body:       []byte(m.ResponseBody),
		Headers:    make(http.Header),
	}

	// Make sure each mocked response header is present in the final response object:
	for header := range m.ResponseHeaders {
		response.Headers.Set(header, m.ResponseHeaders.Get(header))
	}
	return &response, nil
}
