package handler

import "net/http"

// Of wrap h into normal http.HandlerFunc
func Of(h func(*http.Request) *Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := h(r)

		if resp == nil {
			return
		}

		mergeHeader(w.Header(), resp.extraHeader)
		if resp.handler != nil {
			resp.handler(w, r)
		}
	}
}
