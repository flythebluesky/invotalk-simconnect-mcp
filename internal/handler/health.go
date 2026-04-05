package handler

import (
	"encoding/json"
	"net/http"
)

func Health(isConnected func() bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":               "ok",
			"simconnect_connected": isConnected(),
		})
	}
}