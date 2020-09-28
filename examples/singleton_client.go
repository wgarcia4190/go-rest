package examples

import (
	"time"

	"github.com/wgarcia4190/go-rest/gorest"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gorest.Client {
	client := gorest.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		Build()

	return client
}
