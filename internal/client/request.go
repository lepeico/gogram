package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

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
	readers := append([]io.Reader{}, buf)
	w := multipart.NewWriter(buf)
	defer w.Close()

	for k, f := range payload {
		switch val := f.(type) {
		case string:
			w.WriteField(k, val)
		case *os.File:
			mediaHeader := textproto.MIMEHeader{}
			mediaHeader.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v\".", val.Name()))
			mediaHeader.Set("Content-ID", "media")
			mediaHeader.Set("Content-Filename", val.Name())

			mediaPart, _ := w.CreatePart(mediaHeader)
			data, _ := ioutil.ReadAll(val)

			io.Copy(mediaPart, bytes.NewBuffer(data))
		}
	}

	readers = append(readers, bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", w.Boundary())))
	req, err = http.NewRequest("POST", url, ioutil.NopCloser(io.MultiReader(readers...)))
	req.Header.Set("Content-type", w.FormDataContentType())
	return
}
