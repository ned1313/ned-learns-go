package main

import "fmt"

func main() {
	if i :=15; i < 5 {
		fmt.Println("i is less than 5")
	} else if i < 10 {
		fmt.Println("i is less than 10")
	} else {
		fmt.Println("i is at least 10")
	}
	fmt.Println("after the if statements")
}