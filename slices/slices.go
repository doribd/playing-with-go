package main

import "fmt"

func main() {

	// slices is an abstraction of array
	// slices are index based and have a size like arrays yet they can be resized

	var names []string
	var name string
	fmt.Print("Please enter your name:")
	fmt.Scan(&name)
	names = append(names, name)
	fmt.Print("Array size=", len(names))
}
