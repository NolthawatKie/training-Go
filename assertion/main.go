package main

import "fmt"

func main() {

	var a interface{}
	a = 10
	//fmt.Printf("%T %v\n", a, a)

	// if not match type, process is not work
	if i, ok := a.(float32); ok {
		fmt.Println(i)
	}

	var b, c Obj
	b = rectangle{w: 4, l: 4}
	c = triangle{w: 4, h: 4}

	if x, ok := b.(rectangle); ok {
		fmt.Println(x.Area())
	}
	if y, ok := c.(rectangle); ok {
		fmt.Println(y.Area())
	}

}

type Obj interface {
	Area() float64
}

type rectangle struct {
	w, l float64
}

func (rec rectangle) Area() float64 {
	return rec.w * rec.l
}

type triangle struct {
	w, h float64
}

func (rec triangle) Area() float64 {
	return rec.w * rec.h * 0.5
}
