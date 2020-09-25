package handler

import (
	"fmt"
	"github.com/jonccrawley/passhash/utils"
	"net/http"

	"github.com/jonccrawley/passhash/definition"
	"github.com/jonccrawley/passhash/model"
)

// A buffered channel that we can send work requests on.
var WorkQueue = make(chan model.WorkRequest, 1000)

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
		http.Error(w, "password exceeds maximum length is 30 character", http.StatusBadRequest)
		return
	}

	requestId := definition.ExecutionRepo.Increment()

	work := model.WorkRequest{Id: requestId, Password: password}

	WorkQueue <- work
	fmt.Printf("Work request queued for id %v \n",requestId)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%v", requestId)
	return
}
