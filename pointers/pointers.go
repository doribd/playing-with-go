package main

import "fmt"

func main() {
	a := 10
	b := &a
	c := a
	fmt.Println(a, b, *b, c)

	p := new(int)
	fmt.Println(p)
	fmt.Println(*p)

	temp := 10
	fmt.Println(temp)
	setValue(&temp, 20)
	fmt.Println(temp)
}

func setValue(param *int, value int) {
	*param = value
}
