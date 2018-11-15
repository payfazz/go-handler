package defresponse

import (
	"net/http"

	"github.com/payfazz/go-handler"
)

// Redirect as handler.Response
func Redirect(status int, url string) handler.Response {
	return handler.Response{
		Status: status,
		Adapter: func(resp handler.Response) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				handler.MergeHeader(w.Header(), resp.Header)
				http.Redirect(w, r, url, resp.Status)
			}
		},
	}
}
