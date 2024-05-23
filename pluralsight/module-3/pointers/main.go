package main

import "fmt"

func main() {
	s := "Hellow World"

	p := &s

	fmt.Println(p)
	fmt.Println(*p)

	*p = "Hoow Woow"

	fmt.Println(s)

	p = new(string)

	fmt.Println(p, *p)
}