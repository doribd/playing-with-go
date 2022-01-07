package main

import "fmt"

/**
* This is only for my practice, this is not the best practice of how to name variables. Better to chose good name.
 */
func main() {
	var i int8 = 20
	var f float32 = 5.6
	fmt.Println(float32(i) + f)
	fmt.Println(i + int8(f))
	var b byte = 10
	fmt.Println(b)
	var x uint = 5
	fmt.Println(x)

	s := `Hello, 
	Dor!`
	s1 := `good morning`
	s2 := s + s1
	fmt.Println(s2)
	fmt.Println(s)
	eb := s[0]
	fmt.Println(string(eb))
	var r rune = 120
	fmt.Println(r)

	var array [3]int
	array[0] = 1
	array[1] = 2
	array[2] = 3

	fmt.Println(array[2])
}
