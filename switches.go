package main

import (
	"fmt"
	"os"
)

func main() {
	word := os.Args[1]
	// each switch statement can be a variable itslef, we can define variable in the switch to be used
	switch len := len(word); word {
	// case may include more than one statement separated by comma
	case "hi", "hello":
		fmt.Println("Hi to you too")
	case "Who":
		fmt.Println("This is Dor")
	case "what":
		fmt.Println("This is an example in go")
		// this means that also the default will take place
		fallthrough
	// this case will do nothing
	case "kuku":
	default:
		fmt.Println("You need to write something as parameter ", len, "characters")
	}
}
