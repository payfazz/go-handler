package defresponse

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"

	handler "github.com/payfazz/go-handler"
)

// Data as Response.
func Data(status int, contentType string, data []byte) handler.Response {
	if len(data) == 0 {
		return handler.Response{Status: status}
	}

	return handler.Response{
		Status: status,
		Header: http.Header{
			"Content-Type":   []string{contentType},
			"Content-Length": []string{strconv.Itoa(len(data))},
		},
		Body: ioutil.NopCloser(bytes.NewReader(data)),
	}
}
