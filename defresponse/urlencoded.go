package defresponse

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/payfazz/go-handler"
)

// URLEncoded as handler.Response
func URLEncoded(status int, data url.Values) handler.Response {
	str := data.Encode()
	if str == "" {
		return handler.Response{Status: status}
	}

	return handler.Response{
		Status: status,
		Header: http.Header{
			"Content-Type": []string{"application/x-www-form-urlencoded"},
		},
		Body: ioutil.NopCloser(bytes.NewBufferString(str)),
	}
}
