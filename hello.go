package main

import "fmt"

// this is how we define a constant
const learning = "Welcome to my app"

func main() {
	var myName string
	var myNumber int

	myName = "Dor"
	myNumber = 1

	fmt.Println(learning, " This is ", myName, " learning app, number ", myNumber)

}
