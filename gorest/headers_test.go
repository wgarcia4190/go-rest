package gorest

import (
	"fmt"
	"net/http"
	"testing"
)

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
