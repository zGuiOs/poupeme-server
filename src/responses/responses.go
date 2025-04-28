package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON return an response in json
func JSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// Erro return an error in JSON
func Erro(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Err string `json:"erro"`
	}{Err: err.Error()})
}
