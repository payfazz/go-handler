package handler

import "net/http"

// Response representation.
type Response struct {
	status  int
	header  http.Header
	body    []byte
	handler http.HandlerFunc
}

// ResponseBuilder is used to create Response
type ResponseBuilder struct {
	inner Response
}

// NewResponseBuilder return instance of ResponseBuilder
func NewResponseBuilder() *ResponseBuilder {
	return &ResponseBuilder{
		Response{
			status: http.StatusOK,
		},
	}
}

// Build the Responese
func (b *ResponseBuilder) Build() *Response {
	return &b.inner
}

// WithStatus set Response status
func (b *ResponseBuilder) WithStatus(status int) *ResponseBuilder {
	b.inner.status = status
	return b
}

// WithHeader add Response header
func (b *ResponseBuilder) WithHeader(header http.Header) *ResponseBuilder {
	b.inner.header = header
	return b
}

// WithBody set Response body, this body will be closed after the response is written
func (b *ResponseBuilder) WithBody(body []byte) *ResponseBuilder {
	b.inner.body = body
	return b
}

// WithHandler set raw handler to the Response and override the normal behaviour
func (b *ResponseBuilder) WithHandler(f http.HandlerFunc) *ResponseBuilder {
	b.inner.handler = f
	return b
}
