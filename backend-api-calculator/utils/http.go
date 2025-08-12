package utils

import (
	"encoding/json"
	"net/http"
	"log/slog"
)

func WriteJSON(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("Failed to write JSON", "error", err)
	}
}

func ParseJSON(r *http.Request, dst any) error {
	return json.NewDecoder(r.Body).Decode(dst)
}
