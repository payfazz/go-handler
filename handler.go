package handler

import "net/http"

// Response representation.
type Response struct {
	extraHeader http.Header
	handler     http.HandlerFunc
}

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

// New return Response that will be handled by handler.
//
// You should use defresponse package. Use this function if defresponse is not enough for you.
func New(handler http.HandlerFunc) *Response {
	return &Response{handler: handler}
}
