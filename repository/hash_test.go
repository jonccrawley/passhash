package repository

import "testing"

func TestStoreAndRetrieve(t *testing.T) {
	repo := HashRepository{}

	repo.Store(2222222,"ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==")
	result := repo.Get(2222222)

	if result != "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==" {
		t.Errorf("Unable to store hash: got %v want %v",
			result, "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==")
	}

}

func TestRetrieveEmptyString(t *testing.T) {
	repo := HashRepository{}

	result := repo.Get(0000000)

	if result != "" {
		t.Errorf("Unable to store hash: got %v want %v",
			result, "'AN EMPTY STRING'")
	}

}