package gorest

import (
	"net"
	"net/http"
	"time"
)

const (
	defaultMaxIdleConnections = 5
	defaultResponseTimeout    = 5 * time.Second
	defaultConnectionTimeout  = 1 * time.Second
)

type ClientBuilder interface {
	SetHeaders(http.Header) ClientBuilder
	SetConnectionTimeout(time.Duration) ClientBuilder
	SetResponseTimeout(time.Duration) ClientBuilder
	SetMaxIdleConnections(int) ClientBuilder
	DisableTimeouts(bool) ClientBuilder
	Build() Client
}

type clientBuilder struct {
	headers            http.Header
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeout    time.Duration
	disableTimeouts    bool
}

func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}
	return builder
}

func (c *clientBuilder) Build() Client {
	client := httpClient{
		builder: c,
	}
	client.clientOnce.Do(func() {
		client.client = &http.Client{
			Timeout: c.getTimeout(),
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.getMaxIdleConnections(),
				ResponseHeaderTimeout: c.getResponseTimeout(),
				DialContext: (&net.Dialer{
					Timeout: c.getConnectionTimeout(),
				}).DialContext,
			},
		}
	})
	return &client
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}

func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout
	return c
}

func (c *clientBuilder) SetMaxIdleConnections(connections int) ClientBuilder {
	c.maxIdleConnections = connections
	return c
}

func (c *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	c.disableTimeouts = disable
	return c
}

func (c *clientBuilder) getTimeout() time.Duration {
	return c.responseTimeout + c.connectionTimeout
}

func (c *clientBuilder) getMaxIdleConnections() int {
	if c.disableTimeouts {
		return 0
	}

	if c.maxIdleConnections > 0 {
		return c.maxIdleConnections
	}

	return defaultMaxIdleConnections
}

func (c *clientBuilder) getResponseTimeout() time.Duration {
	if c.disableTimeouts {
		return 0
	}

	if c.responseTimeout > 0 {
		return c.responseTimeout
	}

	return defaultResponseTimeout
}

func (c *clientBuilder) getConnectionTimeout() time.Duration {
	if c.disableTimeouts {
		return 0
	}

	if c.connectionTimeout > 0 {
		return c.connectionTimeout
	}

	return defaultConnectionTimeout
}
