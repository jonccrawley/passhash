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