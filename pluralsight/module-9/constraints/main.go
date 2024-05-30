package main

import "fmt"

func main() {
	a1 := []int{1,4,7}
	a2 := []float64{1.356,5.56574,65458}
	a3 := []string{"tom","bomb","calm"}

	s1 := add(a1)
	s2 := add(a2)
	s3 := add(a3)

	fmt.Printf("Array: %v sums to %v\n", a1, s1)
	fmt.Printf("Array: %v sums to %v\n", a2, s2)
	fmt.Printf("Array: %v sums to %v\n", a3, s3)
}

type addable interface {
	int | float64 | string
}

func add[V addable](a []V) V {
	var result V
	for _, v := range a {
		result += v
	}
	return result
}