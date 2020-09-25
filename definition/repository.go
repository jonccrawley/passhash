package definition

import "github.com/jonccrawley/passhash/model"

var (
	HashRepo HashRepository
	StatisticsRepo StatisticsRepository
	ExecutionRepo ExecutionRepository
)

type HashRepository interface {

	Store(id uint64,hash string)
	Get(id uint64) string
}

type StatisticsRepository interface {

	Add(duration uint64)
    Get() model.Statistics
}

type ExecutionRepository interface {

	Increment() uint64
	CurrentCount() uint64
}