package client

import (
	"replication/models"
	"time"
)

type Client interface {
	Set(models.StorageItem) error
	Get(key string) (models.StorageItem, error)
	Delete(key string) error
}

type client struct {
	// current host
	cHost string
	// current port
	cPort string
	// 10 Millisecond is minimum
	timeout time.Duration
}

type ClientOptionsEditor = func(c client) client

func NewClient(opts ...ClientOptionsEditor) Client {
	// Create default client
	c := client{
		timeout: time.Millisecond * 10,
	}

	// appy changes
	for _, editor := range opts {
		c = editor(c)
	}

	return c
}

func WithHost(host string) ClientOptionsEditor {
	return func(c client) client {
		c.cHost = host
		return c
	}
}

func WithPort(port string) ClientOptionsEditor {
	return func(c client) client {
		c.cPort = port
		return c
	}
}

func WithTimeout(timeout time.Duration) ClientOptionsEditor {
	return func(c client) client {
		c.timeout = timeout
		return c
	}
}

func (c client) SetCHost(host string) Client {
	c.cHost = host
	return c
}

func (c client) SetPort(port string) Client {
	c.cPort = port
	return c
}

func (c client) Set(models.StorageItem) error {
	// TODO
	return nil
}

func (c client) Get(key string) (models.StorageItem, error) {
	// TODO
	return models.StorageItem{}, nil
}

func (c client) Delete(key string) error {
	return nil
}
