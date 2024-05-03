package mid

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"github.com/Victor-Acrani/go-api-starterkit/foundation/web"

	"go.uber.org/zap"
)

// Recover is a middleware for recovering.
func Recover(log *zap.SugaredLogger) func(http.Handler) http.Handler {

	m := func(handler http.Handler) http.Handler {

		h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// check panic
			defer func() {
				if rec := recover(); rec != nil {
					// get stack trace
					trace := debug.Stack()
					err := fmt.Errorf("PANIC [%v] TRACE[%s]", rec, string(trace))
					// log error
					log.Errorw("PANIC", "message", err.Error())
					// set error in request context
					web.SetError(r.Context(), err)
				}
			}()

			handler.ServeHTTP(w, r)
		})

		return h
	}

	return m
}
