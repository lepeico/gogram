package client

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const MultipartOffset = 68

type Payload map[string]interface{}

func (p Payload) Set(key string, value interface{}) {
	p[key] = value
}

func (p Payload) HasReader() bool {
	for _, f := range p {
		switch f.(type) {
		case io.Reader:
			return true
		}
	}
	return false
}

func newJSONRequest(url string, payload Payload) (req *http.Request, err error) {
	js, _ := json.Marshal(payload)
	req, err = http.NewRequest("POST", url, bytes.NewBuffer(js))
	req.Header.Set("Content-type", "application/json")
	return
}

func newFormRequest(url string, payload Payload) (req *http.Request, err error) {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	defer w.Close()

	for k, f := range payload {
		switch val := f.(type) {
		case string:
			w.WriteField(k, val)
		case *os.File:
			fw, err := w.CreateFormFile(k, val.Name())
			if err != nil {
				return nil, err
			}
			if _, err = io.Copy(fw, val); err != nil {
				return nil, err
			}
		}
	}

	req, err = http.NewRequest("POST", url, buf)
	req.Header.Set("Content-type", w.FormDataContentType())
	req.ContentLength += MultipartOffset
	return
}
