package main

import "fmt"

func main() {
	p := 2
	double(&p)
	fmt.Println(p)
}

// address
func double(p *int) {
	*p = *p * 2
}
