package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// RespondWithJSON sends response to server in JSON with success status and data.
func RespondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// RespondWithError sends an error response with the error.
func RespondWithError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, err.Error(), status)
	log.Printf(err.Error())
}
