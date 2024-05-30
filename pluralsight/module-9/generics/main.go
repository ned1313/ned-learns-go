package main

import "fmt"

func main() {
	testScores := map[string]float64 {
		"Ron": 87.6,
		"Harry": 105,
		"Hermione": 63.5,
		"Neville": 27,		
	}

	c := clone(testScores)

	fmt.Println(testScores["Harry"], c["Harry"], c)

}

func clone[K comparable,V any](s map[K]V) map[K]V {
	result := make(map[K]V, len(s))
	for i,v := range s {
		result[i] = v
	}
	return result
}