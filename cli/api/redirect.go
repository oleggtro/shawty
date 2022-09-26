package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

type createRedirectRequest struct {
	url string `json:"url"`
}

type createRedirectResponse struct {
	Short string `json:"Short"`
}

type Redirect struct {
	Target string
}

func (c *Client) CreateRedirect(target string) (*Redirect, error) {
	rel := &url.URL{Path: "/sec/redirect"}
	u := c.BaseURL.ResolveReference(rel)

	b := createRedirectRequest{target}

	body, _ := json.Marshal(b)
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	req.Header = map[string][]string{
		"Content-Type":  {"application/json"},
		"Authorization": {viper.GetString("token")},
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var short Redirect

	err = json.NewDecoder(resp.Body).Decode(&short)
	fmt.Println("target: ", target)
	fmt.Printf("short: %v\n", short)
	return &short, err

}
