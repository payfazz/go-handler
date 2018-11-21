package handler

import "net/http"

// MergeHeader merge http.Header src into dst, and return dst.
// it is okay for dst or src to be nil.
func MergeHeader(dst, src http.Header) http.Header {
	if dst == nil {
		dst = make(http.Header)
	}
	for k, v := range src {
		dst[k] = append(dst[k], v...)
	}
	return dst
}

// MergeRespHeader merge header to resp.Header.
// it will return resp.
func MergeRespHeader(header http.Header, resp Response) Response {
	resp.Header = MergeHeader(resp.Header, header)
	return resp
}
