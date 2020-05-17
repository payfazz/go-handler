package handler

import "net/http"

func mergeHeader(dst, src http.Header) http.Header {
	if dst == nil {
		dst = make(http.Header)
	}
	for k, vs := range src {
		for _, v := range vs {
			dst.Add(k, v)
		}
	}
	return dst
}

// MergeHeader is used to add extra header to the response.
//
// see Example for how to use this function.
func MergeHeader(header http.Header, resp Response) Response {
	resp.extraHeader = mergeHeader(resp.extraHeader, header)
	return resp
}
