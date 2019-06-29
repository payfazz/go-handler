package defresponse

import (
	"net/http"
	"strconv"

	handler "github.com/payfazz/go-handler"
)

// Data as Response.
func Data(status int, contentType string, data []byte) *handler.Response {
	if len(data) == 0 {
		return handler.
			NewResponseBuilder().
			WithStatus(status).
			Build()
	}

	return handler.
		NewResponseBuilder().
		WithStatus(status).
		WithHeader(http.Header{
			"Content-Type":   []string{contentType},
			"Content-Length": []string{strconv.Itoa(len(data))},
		}).
		WithBody(data).
		Build()
}
