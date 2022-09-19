package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

type Session struct {
	Token string `json:"token"`
}

type createSessionRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Fetch a new session token
func (c *Client) Login(username string, password string) (*Session, error) {
	rel := &url.URL{Path: "/session"}
	u := c.BaseURL.ResolveReference(rel)

	b := createSessionRequest{username, password}

	body, _ := json.Marshal(b)
	resp, err := http.Post(u.String(), "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var sess Session

	err = json.NewDecoder(resp.Body).Decode(&sess)
	return &sess, err

}