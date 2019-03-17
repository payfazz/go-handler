// Package handler provide new signature for handling http request.
//
// for example usage, see example directory.
package handler

import (
	"io"
	"net/http"
)

// Response representation.
// Response is not reusable, i.e. do not use it in multiple request.
type Response struct {
	// HTTP Response Status Code.
	// if zero-value (0), then http.StatusOK is used by default executor.
	Status int

	// HTTP Header.
	// if zero-value (nil), then no header is writen by default executor.
	Header http.Header

	// The Body, this body is closed when Response.Close is called.
	// if zero-value (nil), then no body is writen by default executor.
	Body io.ReadCloser

	// This Executor is used for executing the Response.
	// if zero-value (nil), then default executor will be used.
	// default executor ignore any error that occurs when write the resp.Body to w.
	Executor func(resp Response, w http.ResponseWriter, r *http.Request)
}

// Close the resource owned by this Response.
func (r *Response) Close() error {
	if r.Body != nil {
		return r.Body.Close()
	}
	return nil
}

// Wrap h into normal http.HandlerFunc, the returned function will close the Response.
func Wrap(h func(*http.Request) Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := h(r)
		defer resp.Close()

		executor := resp.Executor
		if executor == nil {
			executor = defaultExecutor
		}
		executor(resp, w, r)
	}
}

func defaultExecutor(resp Response, w http.ResponseWriter, r *http.Request) {
	MergeHeader(w.Header(), resp.Header)
	if resp.Status != 0 {
		w.WriteHeader(resp.Status)
	}
	if resp.Body != nil {
		io.Copy(w, resp.Body)
	}
}
