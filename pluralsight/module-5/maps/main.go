package main

import "fmt"

func main() {
	var m map[string][]string
	fmt.Println(m)

	m = map[string][]string{
		"coffee": {"Coffee", "Espresso", "Cappucino"},
		"tea": {"Hot Tea","Chai", "Chai Latte"},
	}

	fmt.Println(m["coffee"][0])

	m["other"] = []string{ "Hot Chocolate" }

	fmt.Println(m)

	delete(m, "tea")

	fmt.Println(m)

	v, ok := m["tea"]

	fmt.Println(v,ok)

	m2 := m

	m2["coffee"] = []string{"Coffee", "Espresso", "Cappucino", "Latte"}

	fmt.Println(m)
}