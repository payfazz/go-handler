// Package handler provide new signature for handling http request.
//
// for example usage, see example directory
package handler

import (
	"net/http"
)

// Handler func alias
type Handler func(*http.Request) Response

// Convert Handler to http.Handler
func Convert(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(r).Convert()(w, r)
	}
}
