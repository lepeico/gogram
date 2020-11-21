package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	botFormat  = "%s/bot%s/%s"
	fileFormat = "%s/bot%s/%s"
)

type Client struct {
	Token      string
	APIRoot    string
	IsWebhook  bool
	httpClient *http.Client
}

func (c *Client) Call(method string, payload Payload) (result json.RawMessage, err error) {
	// TODO: Generify Call func to return result as requested type struct
	var response Response
	if c.Token == "" {
		return response.Result, errors.New("No Token provided")
	}

	url := fmt.Sprintf(botFormat, c.APIRoot, c.Token, method)
	var req *http.Request
	if payload.HasReader() {
		req, err = newFormRequest(url, payload)
	} else {
		req, err = newJSONRequest(url, payload)
	}
	if err != nil {
		return response.Result, err
	}

	req.Header.Set("Connection", "keep-alive")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return response.Result, err
	}
	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&response)
	if response.OK == false {
		return response.Result, response
	}

	return response.Result, nil
}

func NewClient(token string) *Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    10 * time.Second,
		DisableCompression: false,
	}

	hc := &http.Client{Transport: tr}

	return &Client{
		Token:      token,
		APIRoot:    "https://api.telegram.org",
		IsWebhook:  false,
		httpClient: hc,
	}
}
