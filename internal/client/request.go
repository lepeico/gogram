package client

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

const MultipartOffset = 68

type Payload map[string]interface{}

func (p Payload) Set(key string, value interface{}) {
	p[key] = value
}

func (p Payload) HasReader() bool {
	for _, payloadValue := range p {
		switch payloadValue.(type) {
		case io.Reader:
			return true
		}
	}
	return false
}

func newJSONRequest(url string, payload Payload) (req *http.Request, err error) {
	js, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err = http.NewRequest("POST", url, bytes.NewBuffer(js))
	req.Header.Set("Content-type", "application/json")
	return
}

func newFormRequest(url string, payload Payload) (req *http.Request, err error) {
	buf := new(bytes.Buffer)
	writter := multipart.NewWriter(buf)
	defer writter.Close()

	for payloadKey, payloadValue := range payload {
		switch value := payloadValue.(type) {
		case string:
			writter.WriteField(payloadKey, value)
		case int:
			v := strconv.Itoa(value)
			writter.WriteField(payloadKey, v)
		case *os.File:
			fw, err := writter.CreateFormFile(payloadKey, value.Name())
			if err != nil {
				return nil, err
			}
			if _, err = io.Copy(fw, value); err != nil {
				return nil, err
			}
		}
	}

	req, err = http.NewRequest("POST", url, buf)
	req.Header.Set("Content-type", writter.FormDataContentType())
	req.ContentLength += MultipartOffset
	return
}
