package requests

import(
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tcnksm/go-httpstat"
)

type RequestHandler struct {
	
}

type RequestHeaders map[string]string
type RequestBody map[string]string

type Response struct {
	StatusCode int
	Stats httpstat.Result
	body []byte
}

func (requestHandler *RequestHandler) Get(url string, headers RequestHeaders) (Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	
	if err != nil {
		return Response{-1, httpstat.Result{}, []byte{}},  err
	}

	return requestHandler.Send(request, headers)
}

func (requestHandler *RequestHandler) Post(targetUrl string, headers RequestHeaders, body RequestBody) (Response, error) {
	jsonValue, err := json.Marshal(body)
	if err != nil {
		return Response{-1, httpstat.Result{}, []byte{}}, err
	}

	request, err := http.NewRequest("POST", targetUrl,strings.NewReader(string(jsonValue)))
	if err != nil {
		return Response{-1, httpstat.Result{}, []byte{}}, err
		
	}
	
	return requestHandler.Send(request, headers)
}

func (requestHandler *RequestHandler) Send(request *http.Request, headers RequestHeaders) (Response, error) {
	for key, value := range headers {
		request.Header.Set(key, value)
	}
	
	var stats httpstat.Result
	context := httpstat.WithHTTPStat(request.Context(), &stats)
	request = request.WithContext(context)
		
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return Response{-1, stats, []byte{}}, err
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	return Response{response.StatusCode, stats, responseBody}, err
}

func (response *Response) BodyAsObject(result interface{}) error {
	reader := bytes.NewReader(response.body)
	err := json.NewDecoder(reader).Decode(result)
	return err
}

func (response *Response) BodyAsString() string {
	return string(response.body)
}

	
