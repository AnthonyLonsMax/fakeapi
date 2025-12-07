package util

import (
	"net/http"
	"strconv"
)

func GetIntFromQuery(r *http.Request, key string, defaultValue int) (int, error) {
	query := r.URL.Query().Get(key)
	if query == "" {
		return defaultValue, nil
	}
	parsed, err := strconv.Atoi(query)
	if err != nil {
		return -1, err
	}
	return parsed, nil
}

func GetStringFromQuery(r *http.Request, key, defaultValue string) string {
	query := r.URL.Query().Get(key)
	if query == "" {
		return defaultValue
	}
	return query
}
