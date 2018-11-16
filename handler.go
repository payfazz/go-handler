// Package handler provide new signature for handling http request.
//
// for example usage, see example directory
package handler

import (
	"net/http"
)

// Handler func alias for processing http request
type Handler func(*http.Request) Response

// From Handler to http.Handler
func From(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Execute(h(r), w, r)
	}
}

// Execute the Response
func Execute(resp Response, w http.ResponseWriter, r *http.Request) {
	adapter := resp.Adapter
	if adapter == nil {
		adapter = defAdapter
	}
	adapter(resp)(w, r)
}
