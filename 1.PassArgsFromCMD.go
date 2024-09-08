package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println()

	var argument []string
	argument = os.Args //return []string and  error, if any

	if len(argument) < 2 {
		fmt.Printf("No passed <args> /n")
		os.Exit(0)
	}
	fmt.Println(argument[1:]) //access passed args

	fmt.Println("")
}
