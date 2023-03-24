package main

import (
	"fmt"
	"time"
)

func main() {
	start1 := time.Now()
	doSomething()
	doSomething()
	doSomething()
	time.Sleep(120 * time.Millisecond)
	fmt.Printf("with no goroutine : %d\n", time.Since(start1))

	start2 := time.Now()
	go doSomething()
	go doSomething()
	go doSomething()
	time.Sleep(120 * time.Millisecond)
	fmt.Printf("with goroutine : %d\n", time.Since(start2))

}

func doSomething() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("do something")
}
