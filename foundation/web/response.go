package web

import (
	"context"
	"fmt"
	"net/http"
)

// Respond converts a Go value to JSON and sends it to the client.
func Respond(ctx context.Context, w http.ResponseWriter, jsonData []byte, statusCode int) {
	SetStatusCode(ctx, statusCode)

	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Content-Length", fmt.Sprintf("%d", len(jsonData)))
	w.WriteHeader(statusCode)
	w.Write(jsonData)
}
