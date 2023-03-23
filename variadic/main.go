package main

import "fmt"

func main() {

	variadicString() //valid
	fmt.Println("---")

	variadicString("a", "b", "c") //valid
	fmt.Println("---")

	s := []string{"e", "f", "g", "h"}
	//how to input slice in variadic func
	// ... = spread operator to spread  each value in func
	variadicString(s...)

}

func variadicString(s ...string) {
	if s == nil {
		fmt.Println("...")
	}
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
