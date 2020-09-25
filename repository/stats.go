package repository

import (
	"github.com/jonccrawley/passhash/model"
	"sync"
)

type StatisticsRepository struct {}

var mutex = sync.Mutex{}
var stats model.Statistics

func (* StatisticsRepository) Add(duration uint64) {

	mutex.Lock()
	stats.TotalDuration += duration
	stats.NumberOfRequests++
	mutex.Unlock()
}

func (* StatisticsRepository) Get() model.Statistics{

	if stats.NumberOfRequests == 0 {
		stats.AvgDuration = 0
		return stats
	}

	stats.AvgDuration = stats.TotalDuration / stats.NumberOfRequests
	return stats
}

