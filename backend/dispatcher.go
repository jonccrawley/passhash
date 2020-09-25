package backend

import (
	"fmt"

	"github.com/jonccrawley/passhash/handler"
	"github.com/jonccrawley/passhash/model"
)

var WorkerQueue chan chan model.WorkRequest

func StartDispatcher(numberOfWorkers int) {

	WorkerQueue = make(chan chan model.WorkRequest, numberOfWorkers)

	//Create Workers
	for i := 0; i < numberOfWorkers; i++ {

		fmt.Println("Starting worker", i+1)
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-handler.WorkQueue:
				fmt.Println("Received work request")
				go func() {
					worker := <-WorkerQueue

					fmt.Println("Dispatching work request")
					worker <- work
				}()
			}
		}
	}()
}
