// Package v1 represents types used by the web application for v1.
package v1

import "fmt"

// NewResponseError contains the info of a http response error.
type ResponseError struct {
	Msg  string
	Code int
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("HTTP error: %s, code: %d", e.Msg, e.Code)
}

// NewResponseError returns NewResponseError.
func NewResponseError(msg string, code int) *ResponseError {
	return &ResponseError{Msg: msg, Code: code}
}
