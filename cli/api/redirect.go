package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

type createRedirectRequest struct {
	Url string `json:"url"`
}

type createRedirectResponse struct {
	Short string `json:"shortlink" binding:"required"`
}

type Redirect struct {
	Target string
}

func (c *Client) CreateRedirect(target string) (*createRedirectResponse, error) {
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

	if resp.StatusCode == 401 {
		return nil, fmt.Errorf("unauthorized")
	}

	if resp.StatusCode == 500 {
		return nil, fmt.Errorf("internal_server_error")
	}

	defer resp.Body.Close()

	var short createRedirectResponse

	// fmt.Printf("resp: %q", resp.Body)

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		panic(err2)
	}

	err3 := json.Unmarshal(body, &short)
	if err3 != nil {
		panic(err3)
	}
	fmt.Println("target: ", target)
	fmt.Printf("short: %v\n", short)
	return &short, err

}
