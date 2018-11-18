package handler

import (
	"io"
	"net/http"
)

// Executor func alias
// It MUST close the resp after the execution is done
type Executor func(resp Response, w http.ResponseWriter, r *http.Request)

// DefaultExecutor .
func DefaultExecutor(resp Response, w http.ResponseWriter, r *http.Request) {
	defer resp.Close()
	MergeHeader(w.Header(), resp.Header)
	if resp.Status == 0 {
		resp.Status = http.StatusOK
	}
	if resp.Status != 0 {
		w.WriteHeader(resp.Status)
	}
	if resp.Body != nil {
		io.Copy(w, resp.Body)
	}
}

// Response representation.
// Response is not reusable, do not use it in multiple request.
// Response MUST be closed by the executor
// Once created, Response should be considered as immutable, i.e. do not change its field
type Response struct {
	// HTTP Response Status Code.
	// if zero-value (0), then http.StatusOK is used by default executor
	Status int

	// HTTP Header.
	// if zero-value (nil), then no header is writen by default executor
	Header http.Header

	// The Body, this body is closed when Response.Close is called.
	// if zero-value (nil), then no body is writen by default executor.
	// this body is closed when the Response.Close is called
	Body io.ReadCloser

	// This Executor is used for executing the Response.
	// if zero-value (nil), then default executor will be used
	Executor Executor
}

// Close the resource owned by this Response
func (r *Response) Close() error {
	if r.Body != nil {
		return r.Body.Close()
	}
	return nil
}
