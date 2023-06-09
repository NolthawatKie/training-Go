package main

import "fmt"

type rectangle struct {
	w, l float64
}

// function with object signature
func (rec rectangle) Area() float64 {
	return rec.w * rec.l
}

func (rec *rectangle) setWidth(w float64) {
	rec.w = w
}

func main() {
	rec := rectangle{
		w: 6,
		l: 5,
	}

	rec.w = 6

	fmt.Println(rec.l)
	fmt.Println(rec.w)
	fmt.Println(rec.Area())
	fmt.Println("----")

	rec.setWidth(65)
	fmt.Println(rec.Area())
}
