package handler_test

import (
	"net/http"

	"github.com/payfazz/go-handler"
	"github.com/payfazz/go-handler/defresponse"
)

func Example() {
	http.HandleFunc("/hello", handler.Of(func(r *http.Request) *handler.Response {
		return defresponse.Text(200, "Hello")
	}))

	http.HandleFunc("/hello-with-header", handler.Of(func(r *http.Request) *handler.Response {
		return handler.MergeHeader(http.Header{
			"Test-Header": {"test header value"},
		},
			defresponse.Text(200, "Hello"),
		)
	}))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
