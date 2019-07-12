package handler

import "net/http"

// Response representation.
type Response struct {
	extraHeader http.Header
	handler     http.HandlerFunc
}

// NewResponse create new Response that will be handled by handler.
//
// You should use defresponse package. Use this function if defresponse is not enough for you
func NewResponse(handler http.HandlerFunc) *Response {
	return &Response{handler: handler}
}
