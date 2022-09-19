package api

import (
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{httpClient: httpClient}
	c.UserAgent = "shawty-cli-v1.alpha"
	c.BaseURL, _ = url.Parse("http://localhost:8080")

	return c
}