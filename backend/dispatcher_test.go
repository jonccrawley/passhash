package backend

import "testing"

func TestStartDispatcher(t *testing.T) {
	StartDispatcher(1)
}

func TestStartDispatcherMultiple(t *testing.T) {
	StartDispatcher(2)
}