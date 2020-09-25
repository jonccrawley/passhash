package utils

import (
	"net/http"
)

func ValidateRequestMethod(expected string, requestMethod string, w http.ResponseWriter) bool {

	if requestMethod != expected {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}
	return true
}