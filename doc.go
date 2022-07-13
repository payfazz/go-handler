/*
Package handler provide new signature for handling http request.

# Motivation

The standard signature for handling http request is:

	func(http.ResponseWriter, *http.Request)

it is not convenience to write branching inside it.

So let use following signature for handling http request

	func h(r *http.Requset) http.HandlerFunc

Consider the following:

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

Now we can write it like this:

	func h(r *http.Requset) http.HandlerFunc {
		if ... {
			return defresponse.Text(500, "some error 1")
		}

		...


		if ... {
			return defresponse.Text(500, "some error 2")
		}

		...

		// will compile error if we forget return
		return defresponse.Text(200, "some data")
	}

Then use Of function to get old signature back

	http.ListenAndServe(":8080", handler.Of(h))
*/
package handler
