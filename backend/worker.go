package backend

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/jonccrawley/passhash/definition"
	"github.com/jonccrawley/passhash/model"
)

// NewWorker creates, and returns a new Worker object. Its only argument
// is a channel that the worker can add itself to whenever it is done its
// work.
func NewWorker(id int, workerQueue chan chan model.WorkRequest) Worker {
	// Create, and return the worker.
	worker := Worker{
		ID:          id,
		Work:        make(chan model.WorkRequest),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool)}

	return worker
}

type Worker struct {
	ID          int
	Work        chan model.WorkRequest
	WorkerQueue chan chan model.WorkRequest
	QuitChan    chan bool
}

// This function "starts" the worker by starting a goroutine, that is
// an infinite "for-select" loop.
func (w *Worker) Start() {
	go func() {
		for {
			// Add ourselves into the worker queue.
			w.WorkerQueue <- w.Work

			select {
				case work := <-w.Work:
					start := time.Now()

					fmt.Printf("worker%d: Processing Password for requests id %v\n", w.ID, work.Id)
					sha512Bytes := sha512.Sum512([]byte(work.Password))

					finalSha := base64.StdEncoding.EncodeToString(sha512Bytes[:])

					duration := time.Since(start).Microseconds()
					definition.StatisticsRepo.Add(uint64(duration))

					fmt.Printf("Timmer: %v\n", duration)

					//TODO: add timer
					definition.HashRepo.Store(work.Id,finalSha)

				case <-w.QuitChan:
					// We have been asked to stop.
					fmt.Printf("worker%d stopping\n", w.ID)
					return
			}
		}
	}()
}

// Stop tells the worker to stop listening for work requests.
//
// Note that the worker will only stop *after* it has finished its work.
func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}