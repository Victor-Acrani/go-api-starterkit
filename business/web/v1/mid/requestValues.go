package mid

import (
	"context"
	"net/http"
	"time"
	"github.com/Victor-Acrani/go-api-starterkit/foundation/web"

	"github.com/google/uuid"
)

// RequestValues is middleware to set all values that are passed through context.
func RequestValues(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// set values for request context
		v := web.Values{
			TraceID: uuid.NewString(),
			Now:     time.Now().UTC(),
		}
		// create new context with values from request context
		ctx := context.WithValue(r.Context(), web.Key, &v)

		// pass context through request
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}
