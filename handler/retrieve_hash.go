package handler

import (
	"fmt"
	"github.com/jonccrawley/passhash/utils"
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"github.com/jonccrawley/passhash/definition"
)

func RetrieveHashHandler(w http.ResponseWriter, r *http.Request) {

	if utils.ValidateRequestMethod("GET",r.Method,w) == false {
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/hash/")
	if id == "" {
		http.Error(w, "You must specify a hash id you are trying to retrieve /hash/:id", http.StatusBadRequest)
		return
	}

	if !isInt(id) {
		http.Error(w, "id must be a number", http.StatusBadRequest)
		return
	}

	intVar, _ := strconv.ParseUint(id, 10, 64)
	hashResult := definition.HashRepo.Get(intVar)

	if hashResult == "" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	fmt.Fprintf(w, "%v", hashResult)
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
