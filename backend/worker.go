package backend

import (
	"crypto/sha512"
	"encoding/base64"
	"log"
	"time"

	"github.com/jonccrawley/passhash/definition"
	"github.com/jonccrawley/passhash/model"
)

func NewWorker(id int, workerQueue chan chan model.WorkRequest) Worker {
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

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerQueue <- w.Work

			select {
				case work := <-w.Work:
					start := time.Now()

					log.Printf("worker%d: Processing Password for requests id %v\n", w.ID, work.Id)
					sha512Bytes := sha512.Sum512([]byte(work.Password))

					finalSha := base64.StdEncoding.EncodeToString(sha512Bytes[:])

					duration := time.Since(start).Microseconds()
					definition.StatisticsRepo.Add(uint64(duration))

					log.Printf("Execution length: %v\n", duration)

					definition.HashRepo.Store(work.Id,finalSha)

				case <-w.QuitChan:
					// We have been asked to stop.
					log.Printf("Worker %d stopping\n", w.ID)
					return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}