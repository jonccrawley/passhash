package backend

import (
	"github.com/jonccrawley/passhash/definition"
	"github.com/jonccrawley/passhash/model"
	"github.com/jonccrawley/passhash/repository"
	"testing"
)

var WorkQueue = make(chan model.WorkRequest, 1000)

func TestWorkQueue(t *testing.T) {

	StartDispatcher(1)
	definition.HashRepo = &repository.HashRepository{}

	work := model.WorkRequest{Id: 1, Password: "crawley"}
	WorkQueue <- work
}
