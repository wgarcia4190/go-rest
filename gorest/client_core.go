package gorest

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	contentType     = "Content-Type"
	jsonContentType = "application/json"
	xmlContentType  = "application/xml"
)

func (c *httpClient) do(context context.Context, method string, url string, headers http.Header, body interface{}) (*Response, error) {
	fullHeaders := c.getRequestHeaders(headers)

	requestBody, bodyErr := c.getRequestBody(fullHeaders.Get(contentType), body)

	if bodyErr != nil {
		return nil, errors.New("unable to marshal the body")
	}

	request, reqErr := http.NewRequestWithContext(context, method, url, bytes.NewBuffer(requestBody))

	if reqErr != nil {
		return nil, errors.New("unable to create a new request")
	}

	request.Header = fullHeaders

	response, err := c.client.Do(request)

	if err != nil {
		return nil, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	finalResponse := Response{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Headers:    response.Header,
		Body:       responseBody,
	}

	return &finalResponse, nil
}

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case jsonContentType:
		return json.Marshal(body)
	case xmlContentType:
		return xml.Marshal(body)

	default:
		return json.Marshal(body)
	}
}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)

	// Add common headers to the request:
	addHeaders(c.builder.Headers, &result)

	// Add custom headers to the request:
	addHeaders(requestHeaders, &result)

	return result
}

func addHeaders(headers http.Header, result *http.Header) {
	for header, value := range headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}
}
