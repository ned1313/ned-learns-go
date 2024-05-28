package main

import "fmt"

func main() {
	i := 1

	for {
		fmt.Println(i)
		i += 1
	    if i >= 10 {
			break
		}	
	}

	j := 0

	for j <= 10 {
		j += 2
		fmt.Println(j)
	}

	for k := 0; k <= 10; k += 1 {
		fmt.Println(k)
	}
}