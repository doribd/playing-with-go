package main

import "fmt"

func main() {

	// here we define a map type
	// var myMap = map[int]int   ==> but this isn't an empty map. To have an empty map, we will need to do as follow:
	// map in go is key,value and retrieve by key
	// in go, we can't mix data types within the value. In this case, i am using int and no more types in the value
	var points = make(map[int]int)

	// key = 1, value = 1
	points[1] = 1
	// key = 2, value = 1
	points[2] = 1
	// key = 3, value = 1
	points[3] = 1

	// this will print 'map[1:1 2:1 3:1]'
	fmt.Println(points)

	// this will print 3
	fmt.Println(len(points))

}
