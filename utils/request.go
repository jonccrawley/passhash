package utils

import (
	"crypto/sha512"
	"encoding/base64"
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

func HashString(value string) (string){

	sha512Bytes := sha512.Sum512([]byte(value))
	return base64.StdEncoding.EncodeToString(sha512Bytes[:])
}