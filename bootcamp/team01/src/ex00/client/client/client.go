package client

import (
	"errors"
	"replication/models"
	"time"
)

const (
	MIN_TIMEOUT     = time.Millisecond * 10
	DEFAULT_TIMEOUT = time.Second * 10
)

var (
	ERR_TIMEOUT_TOO_SHORT = errors.New("Timeout can't be less than 10ms.")
)

// Builder.go
type client struct {
	cHost   string        // current host
	cPort   string        // current port
	timeout time.Duration // Duration limit for every request
}

type ClientOptionsEditor = func(c client) (client, error)

func NewClient(opts ...ClientOptionsEditor) (Client, error) {
	// Create default client
	c := client{
		timeout: DEFAULT_TIMEOUT,
	}
	var err error
	// appy changes
	for _, editor := range opts {
		c, err = editor(c)
		if err != nil {
			return c, err
		}
	}

	return c, err
}

func WithHost(host string) ClientOptionsEditor {
	return func(c client) (client, error) {
		c.cHost = host
		return c, nil
	}
}

func WithPort(port string) ClientOptionsEditor {
	return func(c client) (client, error) {
		c.cPort = port
		return c, nil
	}
}

func WithTimeout(timeout time.Duration) ClientOptionsEditor {
	return func(c client) (client, error) {
		if timeout < MIN_TIMEOUT {
			return c, ERR_TIMEOUT_TOO_SHORT
		}
		c.timeout = timeout
		return c, nil
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

type Client interface {
	Set(models.StorageItem) error
	Get(key string) (models.StorageItem, error)
	Delete(key string) error
}
