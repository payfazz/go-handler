package defresponse

import (
	"encoding/json"
	"net/http"
)

// JSON as Response.
// the error of json.Encoder.Encode is ignored.
func JSON(status int, data interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)

		json.NewEncoder(w).Encode(data)
	}
}

// JSONPretty as Response.
// the error of json.Encoder.Encode is ignored.
func JSONPretty(status int, data interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)

		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		enc.Encode(data)
	}
}
