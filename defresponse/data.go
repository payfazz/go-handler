package defresponse

import (
	"net/http"
	"strconv"

	"github.com/payfazz/go-handler"
)

// Data as Response.
func Data(status int, contentType string, data []byte) handler.Response {
	return handler.New(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		w.Header().Set("Content-Length", strconv.Itoa(len(data)))
		w.WriteHeader(status)

		w.Write(data)
	})
}
