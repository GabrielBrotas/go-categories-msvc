package utils_test

import (
	"testing"

	utils "github.com/GabrielBrotas/go-categories-msvc/pkg/utils"
)

func TestStringToUint(t *testing.T) {
	// Test valid string
	validString := "123"
	expectedUint := uint(123)

	actualUint, err := utils.StringToUint(validString)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if actualUint != expectedUint {
		t.Errorf("Expected %d, got %d", expectedUint, actualUint)
	}

	// Test invalid string
	invalidString := "abc"
	_, err = utils.StringToUint(invalidString)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
