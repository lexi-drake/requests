package requests

import(
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	
	"github.com/tcnksm/go-httpstat"
)

func Head(url string, headers http.Header) (Response, error) {
	return sendWithoutData("HEAD", url, headers)
}
	
func Get(url string, headers http.Header) (Response, error) {
	return sendWithoutData("GET", url, headers)
}

func Delete(url string, headers http.Header) (Response, error) {
	return sendWithoutData("DELETE", url, headers)
}

func sendWithoutData(verb string, url string, headers http.Header) (Response, error) {
	request, err := http.NewRequest(verb, url, nil)
	if err != nil {
		return Response{},  err
	}

	return Send(request, headers)
}

func Post(url string, headers http.Header, body interface{}) (Response, error) {
	return sendWithData("POST", url, headers, body)
}

func Put(url string, headers http.Header, body interface{}) (Response, error) {
	return sendWithData("PUT", url, headers, body)
}

func Patch(url string, headers http.Header, body interface{}) (Response, error) {
	return sendWithData("PATCH", url, headers, body)
}

func sendWithData(verb string, url string, headers http.Header, body interface{}) (Response, error) {
	jsonValue, err := json.Marshal(body)
	if err != nil {
		return Response{}, err
	}

	request, err := http.NewRequest(verb, url, strings.NewReader(string(jsonValue)))
	if err != nil {
		return Response{}, err
		
	}
	
	return Send(request, headers)
}

func Send(request *http.Request, headers http.Header) (Response, error) {
	request.Header = headers;
	
	var stats httpstat.Result
	context := httpstat.WithHTTPStat(request.Context(), &stats)
	request = request.WithContext(context)

	t := time.Now()
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return Response{}, err
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	return Response{response.StatusCode, stats, response.Header, responseBody, t}, err
}


	
