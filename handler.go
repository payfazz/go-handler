// Package handler provide new signature for handling http request.
//
// for example usage, see example directory.
package handler

import (
	"net/http"
)

// Handler func alias for processing http request.
// it implement http.Handler.
type Handler func(*http.Request) Response

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Execute(h(r), w, r)
}

// Execute resp with w and r.
func Execute(resp Response, w http.ResponseWriter, r *http.Request) {
	executor := resp.Executor
	if executor == nil {
		executor = DefaultExecutor
	}
	executor(resp, w, r)
}
