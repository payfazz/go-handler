package defresponse

import (
	"net/url"

	handler "github.com/payfazz/go-handler"
)

// URLEncoded as Response.
func URLEncoded(status int, data url.Values) *handler.Response {
	return Data(status, "application/x-www-form-urlencoded", []byte(data.Encode()))
}
