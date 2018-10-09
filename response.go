package requests

import(
	"bytes"
	"encoding/json"

	"github.com/tcnksm/go-httpstat"
)

type Response struct {
	StatusCode int
	Stats httpstat.Result
	body []byte
}

func (response *Response) BodyAsObject(result interface{}) error {
	reader := bytes.NewReader(response.body)
	err := json.NewDecoder(reader).Decode(result)
	return err
}

func (response *Response) BodyAsString() string {
	return string(response.body)
}
