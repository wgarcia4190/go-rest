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

	"github.com/wgarcia4190/go-rest/core"

	"github.com/wgarcia4190/go-rest/gomime"
)

func (c *httpClient) do(context context.Context, method string, url string, headers http.Header, body interface{}) (*core.Response, error) {
	fullHeaders := c.getRequestHeaders(headers)

	requestBody, bodyErr := c.getRequestBody(fullHeaders.Get(gomime.HeaderContentType), body)

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
	finalResponse := core.Response{
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
	case gomime.ContentTypeJson:
		return json.Marshal(body)
	case gomime.ContentTypeXml:
		return xml.Marshal(body)

	default:
		return json.Marshal(body)
	}
}
