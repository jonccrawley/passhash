package handler

import (
	"fmt"
	"github.com/jonccrawley/passhash/utils"
	"log"
	"net/http"

	"github.com/jonccrawley/passhash/definition"
	"github.com/jonccrawley/passhash/model"
)

var WorkQueue = make(chan model.WorkRequest, 100)

func SaveHashHandler(w http.ResponseWriter, r *http.Request) {

	if utils.ValidateRequestMethod("POST",r.Method,w) == false {
		return
	}

	password := r.PostFormValue("password")

	if password == "" {
		http.Error(w, "You must specify a password.", http.StatusBadRequest)
		return
	}

	if len(password) > 30  {
		http.Error(w, "The password provided exceeds maximum length of 30 character", http.StatusBadRequest)
		return
	}

	requestId := definition.ExecutionRepo.Increment()
	work := model.WorkRequest{Id: requestId, Password: password}

	WorkQueue <- work
	log.Printf("Work request queued for id %v \n",requestId)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%v", requestId)
	return
}

