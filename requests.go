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

func (requestHandler *RequestHandler) Get(url string, headers RequestHeaders, result interface{}) (int, httpstat.Result, error) {
	request, err := http.NewRequest("GET", url, nil)
	
	if err != nil {
		return -1, httpstat.Result{}, err
	}

	return requestHandler.Send(request, headers, result)
}

func (requestHandler *RequestHandler) Post(targetUrl string, headers RequestHeaders, body RequestBody, result interface{}) (int, httpstat.Result, error) {
	jsonValue, err := json.Marshal(body)
	if err != nil {
		return -1, httpstat.Result{}, err
	}

	request, err := http.NewRequest("POST", targetUrl,strings.NewReader(string(jsonValue)))
	if err != nil {
		return -1, httpstat.Result{}, err101
		
	}
	
	return requestHandler.Send(request, headers, result)
}

func (requestHandler *RequestHandler) Send(request *http.Request, headers RequestHeaders, result interface{}) (int, httpstat.Result, error) {
	for key, value := range headers {
		request.Header.Set(key, value)
	}
	
	var stats httpstat.Result
	context := httpstat.WithHTTPStat(request.Context(), &stats)
	request = request.WithContext(context)
		
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return -1, stats, err
	}

	defer response.Body.Close()

	responseBody, _ := ioutil.ReadAll(response.Body)
	json.NewDecoder(responseBody).Decode(result)
	return response.StatusCode, stats, nil
}
