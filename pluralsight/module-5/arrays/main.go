package main

import (
	"fmt"
	"slices"
)

func main(){
  var s []string
  fmt.Println(s)

  s = []string{"Coffee","Tea","Hot Chocolate"}

  for i := 0; i < len(s); i++ {
    fmt.Println(s[i])
  }

  s = append(s, "Espresso")

  fmt.Println(s)

  s2 := s

  s2[0] = "Cafe"

  fmt.Println(s,s2)

  slices.Delete(s, 0, 2)

  fmt.Println(s,s2)
}