package rest

import (
	"encoding/json"
	"net/http"
)

func Encode[T any](w http.ResponseWriter, target T) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(target)
}
