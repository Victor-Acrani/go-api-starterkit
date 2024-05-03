package mid

import (
	"fmt"
	"net/http"
	"time"
	"github.com/Victor-Acrani/go-api-starterkit/foundation/web"

	"go.uber.org/zap"
)

// Logger is a middleware for logging the request status.
func Logger(log *zap.SugaredLogger) func(http.Handler) http.Handler {

	m := func(handler http.Handler) http.Handler {

		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// get values from request context
			v := web.GetValues(r.Context())

			// log request
			path := r.URL.Path
			if r.URL.RawQuery != "" {
				path = fmt.Sprintf("%s?%s", path, r.URL.RawQuery)
			}

			log.Infow("request started", "trace_id", v.TraceID, "method", r.Method, "path", path,
				"remoteaddr", r.RemoteAddr)

			handler.ServeHTTP(w, r)

			log.Infow("request completed", "trace_id", v.TraceID, "method", r.Method, "path", path,
				"remoteaddr", r.RemoteAddr, "statuscode", v.StatusCode, "since", time.Since(v.Now))
		})

		return h
	}

	return m
}
