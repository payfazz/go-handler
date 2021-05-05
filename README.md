# go-handler

[![GoDoc](https://godoc.org/github.com/payfazz/go-handler/v2?status.svg)](https://godoc.org/github.com/payfazz/go-handler/v2)

Package handler provide new signature for handling http request.

see example directory.

stdlib Handler signature is `func(http.ResponseWriter, *http.Request)`, and it is not convenience to write branching inside it.

For example:

```go
func h(w http.ResponseWriter, r *http.Request) {
    if ... {
        http.Error(w, "some error 1", 500)
        // it will be disaster if we forget this return
        return
    }

    ...


    if ... {
        http.Error(w, "some error 2", 500)
        // it will be disaster if we forget this return
        return
    }

    ...
    fmt.Fprintln(w, "some data")
}
```

we can rewrite it like this:

```go
func h(r *http.Requset) http.HandlerFunc {
    if ... {
        return defresponse.Text(500, "some error 1")
    }

    ...


    if ... {
        return defresponse.Text(500, "some error 2")
    }

    ...

    // we can't forget this, because it'll be compile error if there is no `return`
    return defresponse.Text(200, "some data")
}
```
