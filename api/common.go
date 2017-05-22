package api

import (
	"encoding/json"
	"net/http"
)

func respondWithJSON(rw http.ResponseWriter, statusCode int, responseObject interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	encoder := json.NewEncoder(rw)
	encoder.Encode(responseObject)
}
