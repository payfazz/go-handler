package defresponse

import (
	"net/http"
	"net/url"
)

// URLEncoded as Response.
func URLEncoded(status int, data url.Values) http.HandlerFunc {
	return Data(status, "application/x-www-form-urlencoded", []byte(data.Encode()))
}
