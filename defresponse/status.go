package defresponse

import (
	"fmt"
	"net/http"

	"github.com/payfazz/go-handler"
)

// Status as Response.
func Status(status int) handler.Response {
	return Text(status, fmt.Sprintf("%d %s", status, http.StatusText(status)))
}
