package defresponse

import (
	"encoding/json"
	"net/http"

	handler "github.com/payfazz/go-handler"
)

// JSON as Response.
// the error of marshaling json is ignored.
func JSON(status int, data interface{}) handler.Response {
	return handler.Response{
		Status: status,
		Executor: func(resp handler.Response, w http.ResponseWriter, r *http.Request) {
			handler.MergeHeader(w.Header(), resp.Header)
			w.Header()["Content-Type"] = []string{"application/json"}
			if resp.Status != 0 {
				w.WriteHeader(resp.Status)
			}
			json.NewEncoder(w).Encode(data)
		},
	}
}
