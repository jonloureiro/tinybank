package rest

import (
	"encoding/json"
	"net/http"
)

func Decode[T any](r *http.Request, target T) error {
	if err := json.NewDecoder(r.Body).Decode(&target); err != nil {
		return err
	}
	return nil
}
