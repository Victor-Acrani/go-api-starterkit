package mid

import (
	"net/http"

	v1 "github.com/Victor-Acrani/go-api-starterkit/business/web/v1"
	"github.com/Victor-Acrani/go-api-starterkit/foundation/web"
)

// ErrorHandler is a middleware that handles all the errors.
func ErrorHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// pass context through request
		handler.ServeHTTP(w, r)

		// get error from request context if exists
		err := web.GetError(r.Context())
		if err != nil {
			// check if error is type ResponseError
			if he, ok := err.(*v1.ResponseError); ok {
				http.Error(w, he.Msg, he.Code)
				web.SetStatusCode(r.Context(), he.Code)
				return
			}

			// for other errors
			http.Error(w, err.Error(), http.StatusInternalServerError)
			web.SetStatusCode(r.Context(), http.StatusInternalServerError)
		}
	})
}
