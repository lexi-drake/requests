# Requests #

A library for making http requests in go

## Usage ##

Get the library

```bash
go get github.com/lexi-drake/requests
```

Make requests

```go
import(
	"fmt"
	requests "github.com/lexi-drake/requests"
)

// the json tag is required for serialization into a json object
type MyData struct {
     Name string `json:"name"`
     Info int `json:"info"`
}
...

url := "http://localhost:8080"
headers := requests.RequestHeaders{"key": "value"}

body := MyData {"Alexa", 11}
response, err := requests.Post(url, headers, body)
if err != nil {
	// handle error
}

// for html responses
fmt.Println(response.BodyAsString())

// for json responses
err = response.BodyAsObject(yourObject)
```
