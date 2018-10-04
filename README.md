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
	requests "github.com/lexi-drake/requests"
)

...

url := "http://localhost:8080"
r := RequestHandler{}
headers := RequestHeaders{"key": "value"}
body := RequestBody { "foo": "bar", "one": "a"}

status, response, err := r.Post(url, headers, body)
if err != nil {
	// handle error
}
```
