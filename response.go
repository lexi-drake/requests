package requests

import(
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	
	"github.com/tcnksm/go-httpstat"
)

type Response struct {
	StatusCode int
	Stats httpstat.Result
	header RequestHeaders
	body []byte
	Time time.Time
}

func (response *Response) GetHeaderValue(string key) (string, error) {
	if response.header.contains(key) {
		return response.header[key], nil
	}
	return "", errors.New(fmt.Sprintf("header key %s does not exist in response headers", key))
}

func (response *Response) GetHeaders() RequestHeaders {
	return response.headers
}

func (response *Response) BodyAsObject(result interface{}) error {
	reader := bytes.NewReader(response.body)
	err := json.NewDecoder(reader).Decode(result)
	return err
}

func (response *Response) BodyAsString() string {
	return string(response.body)
}
