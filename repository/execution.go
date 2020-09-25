package repository

import (
	"sync/atomic"
)

type ExecutionRepository struct {}

var requestCounter uint64

func (* ExecutionRepository) Increment() uint64{

	atomic.AddUint64(&requestCounter, 1)
	return requestCounter
}

func (* ExecutionRepository) CurrentCount() uint64{

	return requestCounter
}

