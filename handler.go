// Package handler provide new signature for handling http request.
//
// for example usage, see example directory.
package handler

import (
	"io"
	"net/http"
)

// Handler func alias for processing http request.
// it implement http.Handler.
type Handler func(*http.Request) Response

// Executor func alias.
type Executor func(resp Response, w http.ResponseWriter, r *http.Request) error

// Response representation.
// Response is not reusable, i.e. do not use it in multiple request.
// Once created, Response should be considered as immutable, i.e. do not change its field.
type Response struct {
	// HTTP Response Status Code.
	// if zero-value (0), then http.StatusOK is used by DefaultExecutor.
	Status int

	// HTTP Header.
	// if zero-value (nil), then no header is writen by DefaultExecutor.
	Header http.Header

	// The Body, this body is closed when Response.Close is called.
	// if zero-value (nil), then no body is writen by DefaultExecutor.
	Body io.ReadCloser

	// This Executor is used for executing the Response.
	// if zero-value (nil), then DefaultExecutor will be used.
	Executor Executor
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := h(r)
	defer resp.Close()
	Execute(resp, w, r)
}

// Execute resp with w and r.
func Execute(resp Response, w http.ResponseWriter, r *http.Request) {
	executor := resp.Executor
	if executor == nil {
		executor = DefaultExecutor
	}
	executor(resp, w, r)
}

// DefaultExecutor .
func DefaultExecutor(resp Response, w http.ResponseWriter, r *http.Request) error {
	var err error
	MergeHeader(w.Header(), resp.Header)
	if resp.Status == 0 {
		resp.Status = http.StatusOK
	}
	if resp.Status != 0 {
		w.WriteHeader(resp.Status)
	}
	if resp.Body != nil {
		_, err = io.Copy(w, resp.Body)
	}
	return err
}

// Close the resource owned by this Response.
func (r *Response) Close() error {
	if r.Body != nil {
		return r.Body.Close()
	}
	return nil
}
