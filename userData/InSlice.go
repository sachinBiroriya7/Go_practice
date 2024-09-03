package main

import (
	"strconv"
)

var allUserSlice [][]string

func storeInSlice(fname string, lname string, email string, age uint, isMale bool) [][]string {

	var user []string //empty sclice

	ageStr := strconv.FormatUint(uint64(age), 10)
	isMaleStr := strconv.FormatBool(isMale)

	user = append(user, fname, lname, email, ageStr, isMaleStr)

	allUserSlice = append(allUserSlice, user) // When you append user to allUser,
	return allUserSlice                       //you should append user as a whole slice, not unpack it with ....

}
