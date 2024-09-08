package main

import (
	"fmt"
	"os"
)

func main() {
	//os.args return a slice of string []string and error, if any
	var args = os.Args

	//0th element in args would be the excutable file name then the passed args
	fmt.Println(args)

	if len(args) < 2 {
		fmt.Println("No args passed")
		os.Exit(1)
	}

	//accessing all elements from the args passed.
	fmt.Println(args[1:])

}
