package main //start pointer

import "fmt"

func main() {
	var s []int //defaul of slice
	if s == nil {
		fmt.Println("It's is nul")
	}
	//create array with len(cap) compling by Go
	a := [...]int{1, 3, 4, 5}

	//create slice
	s = a[:]

	fmt.Printf("len = %d \ncap =  %d \nslice = %p \npointer %p\n", len(s), cap(s), s, &a)

	fmt.Println("--- Append ---")

	s = append(s, 11, 13)
	fmt.Printf("len = %d \ncap =  %d \nslice = %p \npointer %p\n", len(s), cap(s), s, &a)

}
