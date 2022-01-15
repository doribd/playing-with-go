package main

import "fmt"

// custom type creation, define as global
// struct can be compared with Java Class
type User struct {
	firstName string
	lastName  string
	email     string
	number    int
	isOptIn   bool
}

func main() {

	// map can support only 1 data type, structure can hold mixed types
	var userData = User{
		firstName: "firstName",
		lastName:  "lastName",
		email:     "email",
		number:    1,
		isOptIn:   true,
	}

	fmt.Println(userData.firstName)
	fmt.Println(userData.lastName)
	fmt.Println(userData.email)
	fmt.Println(userData.number)
	fmt.Println(userData.isOptIn)
}
