package main

import "fmt"

// custom type creation, define as global
// struct can be compared with Java Class
type User struct {
	firstName string
	lastName  string
	email     string
	number    int
	isOptIn   bool
}

type Foo struct {
	A int
	B string
}

type Bar struct {
	C string
	F Foo
}

type Baz struct {
	D string
	// embedded struct within the struct, allows direct access to its functions
	Foo
}

func main() {

	// embedding is delegation and not inheritance

	// map can support only 1 data type, structure can hold mixed types
	var userData = User{
		firstName: "firstName",
		lastName:  "lastName",
		email:     "email",
		number:    1,
		isOptIn:   true,
	}

	fmt.Println(userData.firstName)
	fmt.Println(userData.lastName)
	fmt.Println(userData.email)
	fmt.Println(userData.number)
	fmt.Println(userData.isOptIn)

	f := Foo{A: 10, B: "Hello"}
	b1 := Bar{C: "Dor", F: f}
	fmt.Println(b1.F.A)

	b2 := Baz{D: "Dor1", Foo: f}
	fmt.Println(b2.A)
	var f2 Foo = b2.Foo
	fmt.Println(f2)
}
