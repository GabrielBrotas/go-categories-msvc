package error

import (
	"fmt"
)

type ErrorCollection struct {
	Errors []error
}

func NewErrorCollection() *ErrorCollection {
	return &ErrorCollection{
		Errors: []error{},
	}
}

func (ec *ErrorCollection) Add(err error) {
	ec.Errors = append(ec.Errors, err)
}

func (ec *ErrorCollection) HasErrors() bool {
	return len(ec.Errors) > 0
}

func (ec *ErrorCollection) Throw() error {
	if len(ec.Errors) == 0 {
		return nil
	}

	errString := "Errors:\n"
	for _, err := range ec.Errors {
		errString += fmt.Sprintf("- %s\n", err)
	}

	return fmt.Errorf(errString)
}
