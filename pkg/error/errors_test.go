package error_test

import (
	"fmt"
	"testing"

	error "github.com/GabrielBrotas/go-categories-msvc/pkg/error"
)

func TestErrorCollection(t *testing.T) {
	ec := error.NewErrorCollection()

	if ec.Errors == nil {
		t.Errorf("ErrorCollection.New() did not create a valid ErrorCollection instance")
	}

	// Test Add()
	errTest := fmt.Errorf("test error")
	ec.Add(errTest)
	if len(ec.Errors) != 1 || ec.Errors[0] != errTest {
		t.Errorf("ErrorCollection.Add() did not add the error properly")
	}

	// Test HasErrors()
	if ec.HasErrors() == false {
		t.Errorf("ErrorCollection.HasErrors() is invalid")
	}

	// Test Throw()
	err := ec.Throw()
	if err == nil || err.Error() != "Errors:\n- test error\n" {
		t.Errorf("ErrorCollection.Throw() did not return the expected error")
	}
}
