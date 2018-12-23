package defresponse

import (
	"fmt"
	"net/http"

	handler "github.com/payfazz/go-handler"
)

// Status as Response.
func Status(status int) handler.Response {
	return Text(status, fmt.Sprintf("%d %s", status, http.StatusText(status)))
}
