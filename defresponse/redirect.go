package defresponse

import (
	"net/http"
)

// Redirect as Response.
func Redirect(status int, url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, status)
	}
}
