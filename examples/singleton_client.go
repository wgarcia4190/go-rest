package examples

import (
	"net/http"
	"time"

	"github.com/wgarcia4190/go-rest/gomime"
	"github.com/wgarcia4190/go-rest/gorest"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gorest.Client {
	headers := make(http.Header)
	headers.Set(gomime.HeaderContentType, gomime.ContentTypeJson)

	client := gorest.NewBuilder().
		SetHeaders(headers).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		SetUserAgent("test-user-agent").
		//SetBaseUrl("https://api.github.com").
		Build()
	return client
}
