package defresponse

import (
	"net/http"
	"strconv"
)

// Data as Response.
func Data(status int, contentType string, data []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if contentType != "" {
			w.Header().Set("Content-Type", contentType)
		}
		w.Header().Set("Content-Length", strconv.Itoa(len(data)))
		w.WriteHeader(status)

		w.Write(data)
	}
}
