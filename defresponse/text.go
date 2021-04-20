package defresponse

import (
	"github.com/payfazz/go-handler"
)

// Text as Response.
func Text(status int, data string) handler.Response {
	return Data(status, "text/plain; charset=utf-8", []byte(data))
}
