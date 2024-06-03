package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	for msg := range ch1 {
		fmt.Println(msg)
	}
}