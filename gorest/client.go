package gorest

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/wgarcia4190/go-rest/core"
)

type Client interface {
	GetWithContext(context.Context, string, ...http.Header) (*core.Response, error)
	PostWithContext(context.Context, string, interface{}, ...http.Header) (*core.Response, error)
	PutWithContext(context.Context, string, interface{}, ...http.Header) (*core.Response, error)
	PatchWithContext(context.Context, string, interface{}, ...http.Header) (*core.Response, error)
	DeleteWithContext(context.Context, string, ...http.Header) (*core.Response, error)
	OptionsWithContext(context.Context, string, ...http.Header) (*core.Response, error)

	Get(string, ...http.Header) (*core.Response, error)
	Post(string, interface{}, ...http.Header) (*core.Response, error)
	Put(string, interface{}, ...http.Header) (*core.Response, error)
	Patch(string, interface{}, ...http.Header) (*core.Response, error)
	Delete(string, ...http.Header) (*core.Response, error)
	Options(string, ...http.Header) (*core.Response, error)
}

type httpClient struct {
	builder    *clientBuilder
	client     core.HttpClient
	clientOnce sync.Once
}

func (c *httpClient) GetWithContext(context context.Context, url string, headers ...http.Header) (*core.Response, error) {
	return c.do(context, http.MethodGet, getUrl(c.builder.baseUrl, url), getHeaders(headers...), nil)
}
func (c *httpClient) PostWithContext(context context.Context, url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(context, http.MethodPost, getUrl(c.builder.baseUrl, url), getHeaders(headers...), body)
}
func (c *httpClient) PutWithContext(context context.Context, url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(context, http.MethodPut, getUrl(c.builder.baseUrl, url), getHeaders(headers...), body)
}
func (c *httpClient) PatchWithContext(context context.Context, url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.do(context, http.MethodPatch, getUrl(c.builder.baseUrl, url), getHeaders(headers...), body)
}
func (c *httpClient) DeleteWithContext(context context.Context, url string, headers ...http.Header) (*core.Response, error) {
	return c.do(context, http.MethodDelete, getUrl(c.builder.baseUrl, url), getHeaders(headers...), nil)
}
func (c *httpClient) OptionsWithContext(context context.Context, url string, headers ...http.Header) (*core.Response, error) {
	return c.do(context, http.MethodOptions, getUrl(c.builder.baseUrl, url), getHeaders(headers...), nil)
}

func (c *httpClient) Get(url string, headers ...http.Header) (*core.Response, error) {
	return c.GetWithContext(context.Background(), url, headers...)
}
func (c *httpClient) Post(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.PostWithContext(context.Background(), url, body, headers...)
}
func (c *httpClient) Put(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.PutWithContext(context.Background(), url, body, headers...)
}
func (c *httpClient) Patch(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	return c.PatchWithContext(context.Background(), url, body, headers...)
}
func (c *httpClient) Delete(url string, headers ...http.Header) (*core.Response, error) {
	return c.DeleteWithContext(context.Background(), url, headers...)
}
func (c *httpClient) Options(url string, headers ...http.Header) (*core.Response, error) {
	return c.OptionsWithContext(context.Background(), url, headers...)
}

func getUrl(baseUrl string, path string) string {
	return fmt.Sprintf("%s%s", baseUrl, path)
}
