package main

import (
	"fmt"
)

/**
* Testings. We can't declare something without using it.
 */
func main() {
	// direct declare with explicit
	var i int = 10
	fmt.Println(i)

	// infer by declaration
	var b = 20
	fmt.Println(b)

	// declaration and not an assignment
	c := 30
	fmt.Println(c)
}
