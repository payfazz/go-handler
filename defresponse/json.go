package defresponse

import (
	"encoding/json"
	"net/http"

	"github.com/payfazz/go-handler"
)

// JSON as handler.Response
func JSON(status int, data interface{}) handler.Response {
	return handler.Response{
		Status: status,
		Executor: func(resp handler.Response, w http.ResponseWriter, r *http.Request) {
			defer resp.Close()
			handler.MergeHeader(w.Header(), resp.Header)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(resp.Status)
			json.NewEncoder(w).Encode(data)
		},
	}
}
