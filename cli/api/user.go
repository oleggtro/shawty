package api

import (
// "bytes"
// "encoding/json"
// "net/http"
// "net/url"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// func (c *Client) CreateUser(user User) (Session, error) {
// rel := &url.URL{Path: "/signup"}
// u := c.BaseURL.ResolveReference(rel)
// body, _ := json.Marshal(user)
// req, err := http.Post(u.String(), "application/json", bytes.NewBuffer(body))
// if err != nil {
// panic(err)
// }
// }