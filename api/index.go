package api

import (
	"context"
	"net/http"
)

// Index endpoint
func (a *API) Index(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, map[string]string{
		"ok": true,
	})
}
