package defresponse

import (
	"net/http"

	"github.com/payfazz/go-handler"
)

// Redirect as Response.
func Redirect(status int, url string) *handler.Response {
	return handler.NewResponseBuilder().
		WithRawHandler(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, url, status)
		}).
		Build()
}
