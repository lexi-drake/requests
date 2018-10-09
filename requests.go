package requests

import(
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

func (requestHandler *RequestHandler) Head(url string, headers RequestHeaders) (Response, error) {
	return requestHandler.sendWithoutData("HEAD", url, headers)
}
	
func (requestHandler *RequestHandler) Get(url string, headers RequestHeaders) (Response, error) {
	return requestHandler.sendWithoutData("GET", url, headers)
}

func (requestHandler *RequestHandler) Delete(url string, headers RequestHeaders) (Response, error) {
	return requestHandler.sendWithoutData("DELETE", url, headers)
}

func (requestHandler *RequestHandler) sendWithoutData(verb string, url string, headers RequestHeaders) (Response, error) {
	request, err := http.NewRequest(verb, url, nil)
	if err != nil {
		return Response{-1, httpstat.Result{}, []byte{}},  err
	}

	return requestHandler.Send(request, headers)
}

func (requestHandler *RequestHandler) Post(url string, headers RequestHeaders, body RequestBody) (Response, error) {
	return requestHandler.sendWithData("POST", url, headers, body)
}

func (requestHandler *RequestHandler) Put(url string, headers RequestHeaders, body RequestBody) (Response, error) {
	return requestHandler.sendWithData("PUT", url, headers, body)
}

func (requestHandler *RequestHandler) Patch(url string, headers RequestHeaders, body RequestBody) (Response, error) {
	return requestHandler.sendWithData("PATCH", url, headers, body)
}

func (requestHandler *RequestHandler) sendWithData(verb string, url string, headers RequestHeaders, body RequestBody) (Response, error) {
	jsonValue, err := json.Marshal(body)
	if err != nil {
		return Response{-1, httpstat.Result{}, []byte{}}, err
	}

	request, err := http.NewRequest(verb, url, strings.NewReader(string(jsonValue)))
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


	
