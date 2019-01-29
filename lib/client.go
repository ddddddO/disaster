package lib

import (
	"net/http"

	"github.com/pkg/errors"
)

type Client struct {
	url string
}

func NewClient(u string) *Client {
	return &Client{
		url: u,
	}
}

func (c *Client) Get() (*http.Response, error) {
	resp, err := http.Get(c.url)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to GET request")
	}

	return resp, nil
}
