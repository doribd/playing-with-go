package main

import "fmt"

func main() {

	var userName string
	var userMobile string

	// get input from the user
	fmt.Println("Please enter you user name:")
	// we need to use the pointer with scan and not the variable
	fmt.Scan(&userName)
	fmt.Println("Please enter your mobile:")
	// only when using the pointer it will work of course
	fmt.Scan(&userMobile)

	fmt.Printf("Your name is %v and your mobile is %v", userName, userMobile)
}
