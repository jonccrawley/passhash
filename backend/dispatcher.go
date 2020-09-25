package backend

import (
	"log"

	"github.com/jonccrawley/passhash/handler"
	"github.com/jonccrawley/passhash/model"
)

var WorkerQueue chan chan model.WorkRequest

func StartDispatcher(numberOfWorkers int) {

	WorkerQueue = make(chan chan model.WorkRequest, numberOfWorkers)

	//Create Workers
	for i := 0; i < numberOfWorkers; i++ {

		log.Println("Starting worker", i+1)
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-handler.WorkQueue:
				log.Println("Received work request")
				go func() {
					worker := <-WorkerQueue

					log.Println("Dispatching work request")
					worker <- work
				}()
			}
		}
	}()
}
