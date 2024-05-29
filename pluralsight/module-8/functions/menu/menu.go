// Package menu manages menus for a restaurant
package menu

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

var in = bufio.NewReader(os.Stdin)

type menuItem struct {
	name   string
	prices map[string]float64
}

// Adds an item to the menu struct
func Add() {
	fmt.Println("Please enter the name of the new item")
	name, _ := in.ReadString('\n')
	menu = append(menu, menuItem{name: name, prices: make(map[string]float64)})
}

// Prints the menu
func Print() {
	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}
}
