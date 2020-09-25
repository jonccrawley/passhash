package repository

import (
	"testing"
)

func TestIncrement(t *testing.T) {
	repo := ExecutionRepository{}

	repo.Increment()

	if requestCounter != 1 {
		t.Errorf("counter was invalid: got %v want %v",
			requestCounter, 1)
	}
}


func TestCurrentCounter(t *testing.T) {
	repo := ExecutionRepository{}

	requestCounter = 100
	repo.Increment()

	if repo.CurrentCount() != 101 {
		t.Errorf("counter was invalid: got %v want %v",
			requestCounter, 101)
	}
}
