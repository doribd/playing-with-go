package main

/**
* In here https://pkg.go.dev/std we can find all the packages that go ships with.
* We can create custom or use packages of other people
 */
import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"unicode/utf8"
)

func printName(name string) {
	//str := "dor ben dov"
	fmt.Println(name)
	fmt.Println(len(name))
	fmt.Println(utf8.RuneCountInString(name))
}

func printX() {
	// initialize the random seed
	rand.Seed(time.Now().UnixNano())

	// generating random number > 0 and < 10
	first := rand.Intn(10)
	second := rand.Intn(10)

	// we will convert to float64, find the max and then convert back to int
	result := int(math.Max(float64(first), float64(second)))
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(result, "Is bigger")
}

func main() {
	printName("Dor Ben Dov")
	printX()
}
