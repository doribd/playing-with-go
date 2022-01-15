package main

import "fmt"

// accessible not only to main, but all the functions within this package
const minimumNameSize int = 2
const minimumMobileNumber int = 10

var userName string
var userMobile string

func main() {

	// get input from the user
	fmt.Println("Please enter you user name:")
	// we need to use the pointer with scan and not the variable
	fmt.Scan(&userName)
	fmt.Println("Please enter your mobile:")
	// only when using the pointer it will work of course
	fmt.Scan(&userMobile)

	isValidName := len(userName) >= minimumNameSize
	isValidMobile := len(userMobile) >= minimumMobileNumber

	if isValidMobile && isValidName {
		fmt.Printf("Your name is %v and your mobile is %v", userName, userMobile)
	} else {
		fmt.Println("Wrong name or mobile number")
	}
}
