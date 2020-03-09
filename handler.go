package handler

import "net/http"

// Note to anyone who read this sourcecode
// (*Response) and http.HandlerFunc are isomorphic, so it can be converted to each other
// -- *Response        -> http.HandlerFunc : use (*Response).ServeHTTP
// -- http.HandlerFunc -> *Response        : use New

// Response representation.
type Response struct {
	extraHeader http.Header
	handler     http.HandlerFunc
}

// ServeHTTP execute the resp
func (resp *Response) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if resp == nil {
		return
	}

	mergeHeader(w.Header(), resp.extraHeader)

	if resp.handler != nil {
		resp.handler(w, r)
	}
}

// Of wrap h into normal http.HandlerFunc
func Of(h func(*http.Request) *Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(r).ServeHTTP(w, r)
	}
}

// New return Response that will be handled by handler.
//
// You should use defresponse package. Use this function if defresponse is not enough for you.
func New(handler http.HandlerFunc) *Response {
	return &Response{handler: handler}
}

// Execute the resp
//
// DEPRECATED: use (*Response).ServeHTTP instead
func Execute(resp *Response, w http.ResponseWriter, r *http.Request) {
	resp.ServeHTTP(w, r)
}
