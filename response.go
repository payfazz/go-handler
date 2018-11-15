package handler

import (
	"io"
	"net/http"
)

// Response representation
type Response struct {
	Status int
	Header http.Header
	Body   io.ReadCloser

	Adapter ResponseAdapter
}

// WithMergedHeader create new Response with headers are merged
func (r Response) WithMergedHeader(src http.Header) Response {
	r.Header = MergeHeader(r.Header, src)
	return r
}

// ResponseAdapter func alias
type ResponseAdapter func(resp Response) http.HandlerFunc

func defAdapter(resp Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		MergeHeader(w.Header(), resp.Header)
		if resp.Status != 0 {
			w.WriteHeader(resp.Status)
		}
		if resp.Body != nil {
			io.Copy(w, resp.Body)
			resp.Body.Close()
		}
	}
}
