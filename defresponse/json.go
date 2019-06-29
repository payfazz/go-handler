package defresponse

import (
	"encoding/json"
	"net/http"

	handler "github.com/payfazz/go-handler"
)

// JSON as Response.
// the error of json.Encoder.Encode is ignored.
func JSON(status int, data interface{}) *handler.Response {
	return handler.
		NewResponseBuilder().
		WithHandler(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(status)
			json.NewEncoder(w).Encode(data)
		}).
		Build()
}
