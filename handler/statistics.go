package handler

import (
	"encoding/json"
	"fmt"
	"github.com/jonccrawley/passhash/definition"
	"github.com/jonccrawley/passhash/utils"
	"net/http"
)

func StatisticsHandler(w http.ResponseWriter, r *http.Request) {

	if !utils.ValidateRequestMethod("GET",r.Method,w) {
		return
	}

	marshalledJson, _ := json.Marshal(definition.StatisticsRepo.Get())

	w.Header().Set("Content-Type", "text/json")
	fmt.Fprintf(w, "%v", string(marshalledJson))
}