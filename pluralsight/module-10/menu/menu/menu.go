// Package menu manages menus for a restaurant
package menu

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

type menuItem struct {
	name   string
	prices map[string]float64
}
type menu []menuItem

func (m menu) print() {
	for _, item := range m {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}
}

func (m *menu) addItem(name string) error {
	name = strings.TrimSpace(name)
	for _, item := range *m {
		if item.name == name {
			return errors.New("item already exists in menu")
		}
	}
	*m = append(*m, menuItem{name: name, prices: make(map[string]float64)})
	return nil
}

// Adds an item to the menu struct
func Add() error {
	fmt.Println("Please enter the name of the new item")
	name, _ := in.ReadString('\n')
    return data.addItem(name)
}

// Prints the menu
func Print() {
	data.print()
}
