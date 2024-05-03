package web

import (
	"context"
	"time"
)

type ctxKey int

const Key ctxKey = 1

// Values represent state for each request.
type Values struct {
	TraceID    string    // request trace ID
	Now        time.Time // request timestamp
	StatusCode int       // response http status code
	Error      error     // request processing error
}

// GetValues returns the values from the context.
func GetValues(ctx context.Context) *Values {
	v, ok := ctx.Value(Key).(*Values)
	if !ok {
		return &Values{
			TraceID: "00000000-0000-0000-0000-000000000000",
			Now:     time.Now(),
		}
	}

	return v
}

// GetTraceID returns the trace id from the context.
func GetTraceID(ctx context.Context) string {
	v, ok := ctx.Value(Key).(*Values)
	if !ok {
		return "00000000-0000-0000-0000-000000000000"
	}
	return v.TraceID
}

// GetTime returns the time from the context.
func GetTime(ctx context.Context) time.Time {
	v, ok := ctx.Value(Key).(*Values)
	if !ok {
		return time.Now()
	}
	return v.Now
}

// SetStatusCode sets the status code back into the context.
func SetStatusCode(ctx context.Context, statusCode int) {
	v, ok := ctx.Value(Key).(*Values)
	if !ok {
		return
	}

	v.StatusCode = statusCode
}

// SetError sets the error back into the context.
func SetError(ctx context.Context, err error) {
	v, ok := ctx.Value(Key).(*Values)
	if !ok {
		return
	}

	v.Error = err
}

// GetError returns the error from the context.
func GetError(ctx context.Context) error {
	v, ok := ctx.Value(Key).(*Values)
	if !ok {
		return nil
	}

	return v.Error
}
