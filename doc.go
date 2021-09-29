/*

Package handler provide new signature for handling http request.

Motivation

The standard signature for handling http request is:
	func(http.ResponseWriter, *http.Request)
it is not convenience to write branching inside it.

So we create new signature for handling http request
	func h(r *http.Requset) handler.Response

Consider the following:
	func h(w http.ResponseWriter, r *http.Request) {
		if ... {
			http.Error(w, "some error 1", 500)
			return // it will be disaster if we forget this return
		}

		...


		if ... {
			http.Error(w, "some error 2", 500)
			return // it will be disaster if we forget this return
		}

		...

		fmt.Fprintln(w, "some data")
	}

Now we can write it like this:
	func h(r *http.Requset) handler.Response {
		if ... {
			return defresponse.Text(500, "some error 1")
		}

		...


		if ... {
			return defresponse.Text(500, "some error 2")
		}

		...

		return defresponse.Text(200, "some data") // we can't forget this, because it'll be compile error if there is no `return`
	}

Then use Of function to get old signature back
	http.ListenAndServe(":8080", handler.Of(h))


NOTE: v1 is not maintaned anymore, please use v2
*/
package handler
