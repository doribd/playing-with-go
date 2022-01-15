package main

import "fmt"

func main() {

	// arrays in go are in fixed size
	// this array will have the size of 50
	// only variables from the same type can be stored within the array
	var firstName string
	var secondName string
	fmt.Print("Enter first name:")
	fmt.Scan(&firstName)
	fmt.Print("\n")
	fmt.Print("Enter second name:")
	fmt.Scan(&secondName)

	var names = [50]string{}
	names[0] = "Dor"
	names[1] = firstName
	names[2] = secondName

	fmt.Println("Array size=", len(names))

}
