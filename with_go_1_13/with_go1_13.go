package main

import (
	"errors"
	"fmt"
)

type QueryError struct {
	Query string
	Err   error
}

func (e *QueryError) Unwrap() error {
	return e.Err
}

func (e *QueryError) Error() string {
	return e.Query
}

func main() {
	err := fmt.Errorf("%v : %w", "foo", &QueryError{"foo", nil})
	fmt.Println(err.Error())

	if errors.Is(errors.New("foo"), errors.New("foo")) {
		fmt.Println("foo boar")
	}

	var x *QueryError
	if errors.As(err, &x) {
		fmt.Println("moo")
	}

}
