package model

type Statistics struct {

	NumberOfRequests uint64 `json:"total"`
	TotalDuration uint64 `json:"-"`
	AvgDuration uint64 `json:"average"`
}
