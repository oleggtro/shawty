package api

import (
	"net/http"
	"net/url"

	"github.com/spf13/viper"
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
	c.BaseURL, _ = url.Parse(viper.GetString("base_url"))

	return c
}
