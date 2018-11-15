package handler

import "net/http"

// MergeHeader merge http.Header
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
