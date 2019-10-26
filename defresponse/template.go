package defresponse

import (
	"html/template"
	"net/http"

	"github.com/payfazz/go-handler"
)

// HTMLTemplate render template as a response.
func HTMLTemplate(code int, template *template.Template, data interface{}) *handler.Response {
	return handler.New(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(code)
		template.Execute(w, data)
	})
}
