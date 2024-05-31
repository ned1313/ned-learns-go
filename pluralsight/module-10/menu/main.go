package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"menu/menu"
)

var in = bufio.NewReader(os.Stdin)

func main() {

loop:
	for {
		fmt.Println("Please select an option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("q) Quit")
		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			menu.Print()
		case "2":
			err := menu.Add()
			if err != nil {
				fmt.Println(fmt.Errorf("\ninvalid input: %w\n", err))
			}
		case "q":
			break loop
		default:
			fmt.Println("Option not found")
		}
	}

}

