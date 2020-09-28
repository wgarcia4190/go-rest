package gorest

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetRequestBody(t *testing.T) {
	// Initialization:
	client := httpClient{}

	t.Run("NoBodyNilResponse", func(t *testing.T) {
		// Execution
		body, err := client.getRequestBody("", nil)

		// Validation
		if err != nil {
			t.Error("no error expected when passing a nil body")
		}

		if body != nil {
			t.Error("no body expected when passing a nil body")
		}
	})

	t.Run("BodyWithJson", func(t *testing.T) {
		// Execution
		requestBody := []string{"John", "Doe"}

		body, err := client.getRequestBody("application/json", requestBody)

		// Validation
		if err != nil {
			t.Error("no error expected when marshaling slice as json")
		}

		if string(body) != `["John","Doe"]` {
			t.Error("invalid json body obtained")
		}
	})

	t.Run("BodyWithXml", func(t *testing.T) {
		// Execution
		requestBody := []string{"John", "Doe"}

		body, err := client.getRequestBody("application/xml", requestBody)

		// Validation
		if err != nil {
			t.Error("no error expected when marshaling slice as json")
		}

		fmt.Println(string(body))
		if string(body) != `<string>John</string><string>Doe</string>` {
			t.Error("invalid json body obtained")
		}

	})

	t.Run("BodyWithJsonAsDefault", func(t *testing.T) {
		// Execution
		requestBody := []string{"John", "Doe"}

		body, err := client.getRequestBody("", requestBody)

		// Validation
		if err != nil {
			t.Error("no error expected when marshaling slice as json")
		}

		if string(body) != `["John","Doe"]` {
			t.Error("invalid json body obtained")
		}
	})
}

func TestGetRequestHeaders(t *testing.T) {
	// Initialization
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "go-rest-client-agent")
	client.builder = &clientBuilder{
		headers: commonHeaders,
	}

	// Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "QWERTY-789")

	finalHeaders := client.getRequestHeaders(requestHeaders)

	//Validation
	totalHeaders := len(commonHeaders) + len(requestHeaders)
	if len(finalHeaders) != totalHeaders {
		t.Error(fmt.Sprintf("Expecting %d headers", totalHeaders))
	}
	if finalHeaders.Get("X-Request-Id") != "QWERTY-789" {
		t.Error("invalid request id received")
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("invalid content type received")
	}

	if finalHeaders.Get("User-Agent") != "go-rest-client-agent" {
		t.Error("invalid user agent received")
	}
}
