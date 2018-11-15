package defresponse

import (
	"encoding/json"
	"net/http"

	"github.com/payfazz/go-handler"
)

// JSON as handler.Response, it will panic if data is not json.Marshal-able
func JSON(status int, data interface{}) handler.Response {
	return handler.Response{
		Status: status,
		Adapter: func(resp handler.Response) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				handler.MergeHeader(w.Header(), resp.Header)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(resp.Status)
				if err := json.NewEncoder(w).Encode(data); err != nil {
					panic(err)
				}
			}
		},
	}
}
