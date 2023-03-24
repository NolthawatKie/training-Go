package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		start1 := time.Now()
		doSomething()
		doSomething()
		doSomething()
		time.Sleep(120 * time.Millisecond)
		fmt.Printf("with no goroutine : %d\n", time.Since(start1))
	*/

	wg := &sync.WaitGroup{}
	wg.Add(3)

	start2 := time.Now()
	go doSomething(wg)
	go doSomething(wg)
	go doSomething(wg)
	//time.Sleep(120 * time.Millisecond)
	wg.Wait() //Bottle neck
	fmt.Printf("with goroutine : %d\n", time.Since(start2))

}

func doSomething(wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("do something")
	wg.Done()
}
