package repository

import (
	"fmt"
	"sync"
)

type HashRepository struct{}

var hashes sync.Map

func (* HashRepository) Store(id uint64,hash string) {

	hashes.Store(id,hash)
}

func (* HashRepository) Get(id uint64) string{

	result, _ := hashes.Load(id)

	if result == nil {
		result = ""
	}

	return fmt.Sprintf("%v", result)
}