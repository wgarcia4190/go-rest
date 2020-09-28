package gorest

import (
	"net/http"

	"github.com/wgarcia4190/go-rest/gomime"
)

func getHeaders(headers ...http.Header) http.Header {
	if len(headers) > 0 {
		return headers[0]
	}
	return http.Header{}
}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)

	// Add common headers to the request:
	addHeaders(c.builder.headers, &result)

	// Add custom headers to the request:
	addHeaders(requestHeaders, &result)

	// Set User-Agent if it is defined and not there yet:
	if c.builder.userAgent != "" {
		if result.Get(gomime.HeaderUserAgent) != "" {
			return result
		}
		result.Set(gomime.HeaderUserAgent, c.builder.userAgent)
	}
	return result
}

func addHeaders(headers http.Header, result *http.Header) {
	for header, value := range headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}
}
