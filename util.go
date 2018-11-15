package handler

import "net/http"

// MergeHeader merge http.Header src into dst, and return dst
func MergeHeader(dst, src http.Header) http.Header {
	if dst == nil {
		dst = make(http.Header)
	}
	if src != nil {
		for k, v := range src {
			dst[k] = append(dst[k], v...)
		}
	}
	return dst
}
