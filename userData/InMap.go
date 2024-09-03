package main

import (
	"strconv"
)

var allUserMap = make([]map[string]string, 0)

func storeInMap(fname string, lname string, email string, age uint, isMale bool) []map[string]string {
	var UserMap = make(map[string]string, 0)
	UserMap["firstName"] = fname
	UserMap["LastName"] = lname
	UserMap["email"] = email
	UserMap["Age"] = strconv.FormatUint(uint64(age), 10)
	UserMap["Is-Male"] = strconv.FormatBool(isMale)

	allUserMap = append(allUserMap, UserMap)
	return allUserMap

	/*
		allUser = append(allUser, user)
		can't append to a map like you can with slices. Instead, you update the map by assigning values to keys.
	*/

}
