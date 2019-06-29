package handler

import "net/http"

// Of wrap h into normal http.HandlerFunc
func Of(h func(*http.Request) *Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := h(r)

		if resp.handler != nil {
			resp.handler(w, r)
			return
		}

		mergeHeader(w.Header(), resp.header)
		w.WriteHeader(resp.status)

		if resp.body != nil {
			w.Write(resp.body)
		}
	}
}
