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

// MergeRespHeader merge header to resp.Header.
// it will return resp.
func MergeRespHeader(header http.Header, resp *Response) *Response {
	resp.header = mergeHeader(resp.header, header)
	return resp
}
