package main

import "fmt"

const maxCounter int = 10

func main() {
	// simple count up to 7 with a simple break
	counter := 0
	for counter < maxCounter {
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
		if counter > maxCounter {
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
