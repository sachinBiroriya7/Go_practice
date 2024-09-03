package main

type Person struct {
	firstName string
	lastName  string
	email     string
	age       uint
	isMale    bool
}

var allUserStruct = make([]Person, 0)

func storeInStruct(fname string, lname string, email string, age uint, isMale bool) []Person {
	var user Person
	user.firstName = fname
	user.lastName = lname
	user.email = email
	user.age = age
	user.isMale = isMale

	allUserStruct = append(allUserStruct, user)
	return allUserStruct

}
