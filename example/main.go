package main

import (
	"net/http"

	handler "github.com/payfazz/go-handler"
	"github.com/payfazz/go-handler/defresponse"
)

func main() {
	h := http.NewServeMux()
	h.Handle("/", handler.Wrap(root))
	h.Handle("/test", handler.Wrap(test))

	if err := http.ListenAndServe(":8080", h); err != nil {
		panic(err)
	}
}

func root(r *http.Request) handler.Response {
	customHeader := make(http.Header)
	customHeader.Set("X-Asdf", "lala")
	customHeader.Set("X-Lala", "asdf")

	return handler.MergeRespHeader(
		customHeader,
		defresponse.JSON(http.StatusOK, struct {
			A string `json:"a"`
			B string `json:"b"`
		}{
			A: "Hello world",
			B: r.URL.EscapedPath(),
		}),
	)
}

func test(r *http.Request) handler.Response {
	if r.Header.Get("Authorization") == "" {
		return defresponse.Status(http.StatusUnauthorized)
	}

	if r.Method != http.MethodGet {
		return defresponse.Status(http.StatusMethodNotAllowed)
	}

	return defresponse.Text(http.StatusOK, "some data")
}
