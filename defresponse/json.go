package defresponse

import (
	"encoding/json"
	"net/http"

	"github.com/payfazz/go-handler"
)

// JSON as Response.
// the error of marshaling json is ignored.
func JSON(status int, data interface{}) handler.Response {
	return handler.Response{
		Status: status,
		Executor: func(resp handler.Response, w http.ResponseWriter, r *http.Request) error {
			handler.MergeHeader(w.Header(), resp.Header)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(resp.Status)
			return json.NewEncoder(w).Encode(data)
		},
	}
}
