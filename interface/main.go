package main

import "fmt"

func main() {

	//empty interface
	var a interface{}

	// %T : type
	// %v : value
	a = 10
	fmt.Printf("%T %v\n", a, a)

	a = "ten"
	fmt.Printf("%T %v\n", a, a)

	a = true
	fmt.Printf("%T %v\n", a, a)

	a = func() string { return "Hello" }
	fmt.Printf("%T %v\n", a, a)

}
