package main

import(
//	"bytes"
	"encoding/json"
//	"fmt"
//	"io"
	"io/ioutil"
	"log"
	"net/http"
//	"net/http/httputil"
	"strings"
)

type RequestHandler struct {
	
}

type RequestHeaders map[string]string
type RequestBody map[string]string

func (requestHandler *RequestHandler) Get(url string, headers RequestHeaders, result interface{}) (int, string, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return -1, "", err
	}
	
	return requestHandler.Send(request, headers)
}

func (requestHandler *RequestHandler) Post(targetUrl string, headers RequestHeaders, body RequestBody ) (int, string, error) {
	jsonValue, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", targetUrl,strings.NewReader(string(jsonValue)))
	if err != nil {
		return -1, "", err
	}
	
	return requestHandler.Send(request, headers)
}

func (requestHandler *RequestHandler) Send(request *http.Request, headers RequestHeaders) (int, string, error) {
	for key, value := range headers {
		request.Header.Set(key, value)
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return -1, "", err
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response.StatusCode, "", err
	}
	
	body := string(responseBody)
	return response.StatusCode, body, nil
}
