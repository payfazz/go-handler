package handler_test

import (
	"net/http"

	"github.com/payfazz/go-handler"
	"github.com/payfazz/go-handler/defresponse"
)

func Example() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handler.Of(func(r *http.Request) *handler.Response {
		return defresponse.Text(200, "Hello")
	}))
	panic(http.ListenAndServe(":8080", mux))
}
