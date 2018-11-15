package defresponse

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/payfazz/go-handler"
)

// Text as handler.Response
func Text(status int, data string) handler.Response {
	return handler.Response{
		Status: status,
		Header: http.Header{
			"Content-Type": []string{"text/plain; charset=utf-8"},
		},
		Body: ioutil.NopCloser(bytes.NewBufferString(data)),
	}
}
