package v1

import (
	"encoding/json"
	"net/http"

	"github.com/Victor-Acrani/go-api-starterkit/foundation/web"
)

// healthCheck is a handle fucntion for checking is the server is up.
func NotFound(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Status  string `json:"status"`
	}{
		Status:  "404 not found",
	}

	jsonData, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	web.Respond(r.Context(), w, jsonData, http.StatusNotFound)
}
