package main

import "fmt"

const A = 50

func main() {
	const a = 42
	var i int = a
	var I int = A


	println(i)
	println(I)

	const c = iota
	fmt.Println(c)

	const (
		d = 2*5
		e
		f = iota
		g
		h = 10 * iota
	)

	fmt.Println(d,e,f,g,h)
}