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
	fileFormat = "%s/file/bot%s/%s"
)

type Client struct {
	Token      string
	APIRoot    string
	IsWebhook  bool
	httpClient *http.Client
}

func (c *Client) Call(method string, payload Payload, ret interface{}) (err error) {
	// TODO: Generify Call func to return result as requested type struct
	if c.Token == "" {
		return errors.New("No Token provided")
	}

	url := fmt.Sprintf(botFormat, c.APIRoot, c.Token, method)
	var req *http.Request
	if payload.HasReader() {
		req, err = newFormRequest(url, payload)
	} else {
		req, err = newJSONRequest(url, payload)
	}
	if err != nil {
		return err
	}

	req.Header.Set("Connection", "keep-alive")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var response Response
	json.NewDecoder(res.Body).Decode(&response)
	if response.OK == false {
		return response
	}

	json.Unmarshal(response.Result, ret)

	return nil
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
