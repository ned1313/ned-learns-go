package main

import "fmt"

func main() {
	a, b := 5, 2
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	fmt.Println(float32(a) / float32(b))

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)

	var (
		c int = 6
		d string = "hi"
	)
	fmt.Println(c)
	fmt.Println(d)
}