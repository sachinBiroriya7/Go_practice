package main

import "fmt"

func GetUserInfo() (string, string, string, uint, bool) {

	var fname string
	var lname string
	var email string
	var age uint
	var isMale bool

	fmt.Println("enter users first name")
	fmt.Scan(&fname)

	fmt.Println("enter users last name")
	fmt.Scan(&lname)

	fmt.Println("enter users email")
	fmt.Scan(&email)

	fmt.Println("enter users Age")
	fmt.Scan(&age)

	fmt.Println("user isMale? [Y/N]")
	var input string
	fmt.Scan(&input)

	if input == "Y" || input == "y" {
		isMale = true
	} else {
		isMale = false
	}

	return fname, lname, email, age, isMale
}
