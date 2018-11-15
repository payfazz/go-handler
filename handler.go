// Package handler provide new signature for handling http request.
//
// for example usage, see example directory
package handler

import (
	"net/http"
)

// Handler func alias
type Handler func(*http.Request) Response

// From convert Handler to http.Handler
func From(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := h(r)
		if resp.Adapter != nil {
			resp.Adapter(resp)(w, r)
		} else {
			defAdapter(resp)(w, r)
		}
	}
}
