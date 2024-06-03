package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	rounds := 100
	wg.Add(rounds)
	for i := 0; i < rounds; i++ {
		num := strconv.Itoa(i)
	go func(){ 
		buzz(num)
		wg.Done()
	}()
	}

	wg.Wait()

}

func buzz(text string) {
  fmt.Println("zzz",text,"zzz")
}