package utils

import (
	"net/http/httptest"
	"testing"
)

func TestValidateRequestMethod(t *testing.T) {
	w := httptest.NewRecorder()
	isValid := ValidateRequestMethod("GET","GET",w)

	if !isValid {
		t.Errorf("Expected request method validation to be true ")
	}
}

func TestValidateRequestMethodFailed(t *testing.T) {
	w := httptest.NewRecorder()
	isValid := ValidateRequestMethod("GET","POST",w)

	if isValid {
		t.Errorf("Expected request method validation to be false ")
	}
}

func TestHash(t *testing.T) {
	value := HashString("angryMonkey")

	if value != "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==" {
		t.Errorf("Unable to validate hashed password")
	}
}