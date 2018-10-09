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

...

url := "http://localhost:8080"
r := RequestHandler{}
headers := RequestHeaders{"key": "value"}
body := RequestBody { "foo": "bar", "one": "a"}

response, err := r.Post(url, headers, body)
if err != nil {
	// handle error
}

// for html responses
fmt.Println(response.BodyAsString())

// for json responses
err = response.BodyAsObject(yourObject)
```
