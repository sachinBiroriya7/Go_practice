package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("this program demonstrate how we can add users data in different collections")
	fmt.Println("i.e. slices Maps and Structs")

	flag := true
	for flag {
		fname, lname, email, age, isMale := GetUserInfo()

		inSlice := storeInSlice(fname, lname, email, age, isMale)

		inMap := storeInMap(fname, lname, email, age, isMale)

		inStruct := storeInStruct(fname, lname, email, age, isMale)

		flag = addUserAgain()

		fmt.Println(inSlice)
		fmt.Println("******************")
		fmt.Println(inMap)
		fmt.Println("******************")
		fmt.Println(inStruct)

	}
}

func addUserAgain() bool {
	fmt.Println("add another user? {Y/N}")
	var again string
	fmt.Scan(&again)
	again = strings.ToLower(again)

	return again == "y"

	/*if again == "y" {
		return true
	}
	return false*/
}
