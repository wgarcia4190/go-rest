package gorest

import (
	"context"
	"net/http"
	"sync"
)

type Client interface {
	GetWithContext(context.Context, string, http.Header) (*Response, error)
	PostWithContext(context.Context, string, http.Header, interface{}) (*Response, error)
	PutWithContext(context.Context, string, http.Header, interface{}) (*Response, error)
	PatchWithContext(context.Context, string, http.Header, interface{}) (*Response, error)
	DeleteWithContext(context.Context, string, http.Header) (*Response, error)
	OptionsWithContext(context.Context, string, http.Header) (*Response, error)

	Get(string, http.Header) (*Response, error)
	Post(string, http.Header, interface{}) (*Response, error)
	Put(string, http.Header, interface{}) (*Response, error)
	Patch(string, http.Header, interface{}) (*Response, error)
	Delete(string, http.Header) (*Response, error)
	Options(string, http.Header) (*Response, error)
}

type httpClient struct {
	builder    *clientBuilder
	client     *http.Client
	clientOnce sync.Once
}

func (c *httpClient) GetWithContext(context context.Context, url string, headers http.Header) (*Response, error) {
	return c.do(context, http.MethodGet, url, headers, nil)
}
func (c *httpClient) PostWithContext(context context.Context, url string, headers http.Header, body interface{}) (*Response, error) {
	return c.do(context, http.MethodPost, url, headers, body)
}
func (c *httpClient) PutWithContext(context context.Context, url string, headers http.Header, body interface{}) (*Response, error) {
	return c.do(context, http.MethodPut, url, headers, body)
}
func (c *httpClient) PatchWithContext(context context.Context, url string, headers http.Header, body interface{}) (*Response, error) {
	return c.do(context, http.MethodPatch, url, headers, body)
}
func (c *httpClient) DeleteWithContext(context context.Context, url string, headers http.Header) (*Response, error) {
	return c.do(context, http.MethodDelete, url, headers, nil)
}
func (c *httpClient) OptionsWithContext(context context.Context, url string, headers http.Header) (*Response, error) {
	return c.do(context, http.MethodOptions, url, headers, nil)
}

func (c *httpClient) Get(url string, headers http.Header) (*Response, error) {
	return c.GetWithContext(context.Background(), url, headers)
}
func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.PostWithContext(context.Background(), url, headers, body)
}
func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.PutWithContext(context.Background(), url, headers, body)
}
func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.PatchWithContext(context.Background(), url, headers, body)
}
func (c *httpClient) Delete(url string, headers http.Header) (*Response, error) {
	return c.DeleteWithContext(context.Background(), url, headers)
}
func (c *httpClient) Options(url string, headers http.Header) (*Response, error) {
	return c.OptionsWithContext(context.Background(), url, headers)
}
