package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("this is an error")
	fmt.Println(err)

	err2 := fmt.Errorf("this error wraps the first one: %w", err)
	fmt.Println(err2)

	err3 := errors.Unwrap(err2)
	fmt.Println(err3)
}