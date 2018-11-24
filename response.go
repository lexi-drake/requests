package requests

import(
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	
	"github.com/tcnksm/go-httpstat"
)

type Response struct {
	StatusCode int
	Stats httpstat.Result
	header http.Header
	body []byte
	Time time.Time
}

func containsKey(headers http.Header, key string) bool {
	return headers.Get(key) != ""
}

func (response *Response) GetHeaderValue(key string) ([]string, error) {
	if containsKey(response.header, key) {
		return response.header[key], nil
	}
	return []string{}, errors.New(fmt.Sprintf("header key %s does not exist in response headers", key))
}

func (response *Response) GetHeader() http.Header {
	return response.header
}

func (response *Response) BodyAsObject(result interface{}) error {
	reader := bytes.NewReader(response.body)
	err := json.NewDecoder(reader).Decode(result)
	return err
}

func (response *Response) BodyAsString() string {
	return string(response.body)
}
