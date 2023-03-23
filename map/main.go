package main

import "fmt"

func main() {
	/*
		//create map
		var m map[string]string

		//make a map
		m = make(map[string]string)
	*/

	/*
		//one line to create map
		m := map[string]string{}
		m["a"] = "apple"
		m["b"] = "banana"
	*/

	m := map[string]string{
		"a": "apple",
		"b": "banana",
	}

	if m == nil {
		fmt.Println("It's null")
	}

	fmt.Println(m)

}
