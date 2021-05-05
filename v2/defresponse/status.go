package defresponse

import (
	"fmt"
	"net/http"
)

// Status as Response.
func Status(status int) http.HandlerFunc {
	return Text(status, fmt.Sprintf("%d %s", status, http.StatusText(status)))
}
