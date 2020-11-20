package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Payload map[string]interface{}

func (p Payload) Set(key string, value interface{}) {
	p[key] = value
}

func newJSONRequest(payload Payload, url string) (req *http.Request, err error) {
	js, _ := json.Marshal(payload)
	req, err = http.NewRequest("POST", url, bytes.NewBuffer(js))
	req.Header.Set("Content-type", "application/json")
	return req, err
}

// func newFormRequest(payload Payload, url string) (req *http.Request, err error) {
// 	var b bytes.Buffer
// 	w := multipart.NewWriter(&b)
// 	defer w.Close()
// 	for key, r := range payload {
// 		var fw io.Writer
// 		if x, ok := r.(io.Closer); ok {
// 			defer x.Close()
// 		}
// 		if x, ok := r.(*os.File); ok {
// 			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
// 				return
// 			}
// 		} else {
// 			if fw, err = w.CreateFormField(key); err != nil {
// 				return
// 			}
// 		}
// 		if _, err = io.Copy(fw, r); err != nil {
// 			return err
// 		}
// 	}

// 	req, err = http.NewRequest("POST", url, &b)
// 	req.Header.Set("Content-type", w.FormDataContentType())
// 	return req, err
// }
