package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("What would you like me to scream?")
	in := bufio.NewReader(os.Stdin)

	response, _ := in.ReadString('\n')

	response = strings.TrimSpace(response)
	
	fmt.Println(strings.ToUpper(response) + "!")
}