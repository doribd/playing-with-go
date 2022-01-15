package main

import "fmt"

func main() {
	// simple count up to 7 with a simple break
	counter := 0
	for counter < 10 {
		if counter > 7 {
			break
		}
		fmt.Println(counter)
		counter = counter + 1
	}
	fmt.Println("***************************")
	// counting up to 10
	counter = 0
	for {
		if counter > 10 {
			break
		}
		fmt.Println(counter)
		counter = counter + 1
	}

	// another for loop, now going over the runes
	s := "Hi, This is Dor"
	for x, y := range s {
		fmt.Println(x, y, string(y))
	}
}
