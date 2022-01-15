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
	fmt.Println("Array size=", len(names))

	// example of a loop over the slice using the range
	for index := range names {
		fmt.Println("Name [", index, "]=", names[index])
	}
}
