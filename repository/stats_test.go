package repository

import "testing"

func TestIncrementAndGet(t *testing.T) {
	repo := StatisticsRepository{}

	repo.Add(1234)
	result := repo.Get()

	if result.AvgDuration != 1234 {
		t.Errorf("Invalid average duration: got %v want %v",
			result.AvgDuration, 1234)
	}

	if result.TotalDuration != 1234 {
		t.Errorf("Invalid total duration: got %v want %v",
			result.TotalDuration, 1234)
	}

	if result.NumberOfRequests != 1 {
		t.Errorf("Invalid number fo requests: got %v want %v",
			result.NumberOfRequests, 1)
	}

}

func TestDivideByZero(t *testing.T) {
	stats.NumberOfRequests = 0
	repo := StatisticsRepository{}

	result := repo.Get()

	if result.AvgDuration != 0 {
		t.Errorf("Invalid average duration: got %v want %v",
			result.AvgDuration, 0)
	}

}