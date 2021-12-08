package handler

import (
	"net/http"
)

// Of wrap h into normal http.HandlerFunc
func Of(h func(*http.Request) http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(r)(w, r)
	}
}

// MergeHeader is used to add extra header to the resp.
func MergeHeader(header http.Header, resp http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for k, vs := range header {
			for _, v := range vs {
				w.Header().Add(k, v)
			}
		}
		resp(w, r)
	}
}
