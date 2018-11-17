package handler

import (
	"io"
	"net/http"
)

// ResponseAdapter func alias
type ResponseAdapter func(resp Response) http.HandlerFunc

// Response representation.
// Response is not reusable, do not use it in multiple request,
// because Body are already closed at the end of first execution
type Response struct {
	// HTTP Response Status Code
	Status int

	// Header
	Header http.Header

	// This body will be closed at the end of the handler
	// wrap with ioutil.NopCloser if you want to keep it open
	Body io.ReadCloser

	// This Adapter is used for converting this Response
	// into http.HandlerFunc, if nil it will use default Adapter
	Adapter ResponseAdapter
}

// WithMergedHeader create new Response with headers are merged
func (r Response) WithMergedHeader(src http.Header) Response {
	r.Header = MergeHeader(r.Header, src)
	return r
}

func defAdapter(resp Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		MergeHeader(w.Header(), resp.Header)
		if resp.Status == 0 {
			resp.Status = http.StatusOK
		}
		w.WriteHeader(resp.Status)
		if resp.Body != nil {
			io.Copy(w, resp.Body)
			resp.Body.Close()
		}
	}
}
