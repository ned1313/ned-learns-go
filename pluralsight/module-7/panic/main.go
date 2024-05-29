package main

import "fmt"

func main() {
  fmt.Println("Before the divide function")
  result := divide(8,8)
  fmt.Println("Result is",result)

  result = divide(8,0)
  fmt.Println("Result is", result)
}

func divide(dividend, divisor int) int {
	defer func(){
		if msg := recover(); msg != nil {
		fmt.Println(msg)
		}
	}()
	return dividend/ divisor
}