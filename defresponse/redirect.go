package defresponse

import (
	"net/http"

	handler "github.com/payfazz/go-handler"
)

// Redirect as Response.
func Redirect(status int, url string) *handler.Response {
	return handler.NewResponseBuilder().
		WithHandler(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, url, status)
		}).
		Build()
}
