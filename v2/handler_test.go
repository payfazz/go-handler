package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/payfazz/go-handler/v2"
	"github.com/payfazz/go-handler/v2/defresponse"
)

func ExampleOf() {
	mux := handler.Of(func(r *http.Request) http.HandlerFunc {
		return defresponse.Text(200, "Hello")
	})

	req := httptest.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	mux.ServeHTTP(res, req)

	fmt.Println(res.Body.String())
	// Output: Hello
}

func ExampleMergeHeader() {
	mux := handler.Of(func(r *http.Request) http.HandlerFunc {
		return handler.MergeHeader(http.Header{
			"Test-Header": {"test header value"},
		},
			defresponse.Text(200, "Hello"),
		)
	})

	req := httptest.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	mux(res, req)

	fmt.Println(res.Header().Get("Test-Header"))
	// Output: test header value
}
