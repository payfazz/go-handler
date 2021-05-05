package defresponse

import "net/http"

// Text as Response.
func Text(status int, data string) http.HandlerFunc {
	return Data(status, "text/plain; charset=utf-8", []byte(data))
}
