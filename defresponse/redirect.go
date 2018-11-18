package defresponse

import (
	"net/http"

	"github.com/payfazz/go-handler"
)

// Redirect as handler.Response
func Redirect(status int, url string) handler.Response {
	return handler.Response{
		Status: status,
		Executor: func(resp handler.Response, w http.ResponseWriter, r *http.Request) {
			defer resp.Close()
			handler.MergeHeader(w.Header(), resp.Header)
			http.Redirect(w, r, url, resp.Status)
		},
	}
}
