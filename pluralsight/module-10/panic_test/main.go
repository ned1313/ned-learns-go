package main

import (
	"fmt"
	"errors"
)

func main() {
  fmt.Println("Before the divide function")
  result, err := divide(8,8)
  if err != nil {
	fmt.Println(err)
	return
  }
  fmt.Println("Result is",result)

  result, err = divide(8,0)
  if err != nil {
	fmt.Println(err)
	return
  }
  fmt.Println("Result is", result)
}

func divide(dividend, divisor int) (result int, err error) {
	defer func() {
		if msg := recover(); msg != nil {
			result = 0
			err = fmt.Errorf("%v", msg)
		}
	}()
	if divisor == 0 {
		return 0, errors.New("divisor cannot be zero")
	}
	return dividend/ divisor, nil
}