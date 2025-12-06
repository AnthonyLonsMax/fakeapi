package util

import (
	"encoding/json"
	"net/http"
)

func SendJson(data any, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

func PartialUpdate(src map[string]any, dest map[string]any) {
	for k := range src {
		if _, ok := dest[k]; ok {
			dest[k] = src[k]
		}
	}
}
