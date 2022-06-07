package web

import (
	"fmt"
)

type error interface {
	Error() string
}

type CustomError struct {
	Status int
	Err    error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%d - %v", e.Status, e.Err)
}
