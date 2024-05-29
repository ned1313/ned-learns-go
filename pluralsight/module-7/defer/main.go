package main

import "fmt"

func main() {
	fmt.Println("main 1")

	defer fmt.Println("defer 1")

	fmt.Println("main 2")

	defer fmt.Println("defer 2")

	hey := testing()

	fmt.Println(hey)
}

func testing() string {
	
	defer fmt.Println("testing defer")

	return "test"
}