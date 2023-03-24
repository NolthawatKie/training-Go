package main

import (
	"fmt"
	"sync"
	"time"
)

var i int
var mux sync.Mutex

type thing struct {
	i   int
	mux sync.Mutex
}

func main() {
	/*
		start1 := time.Now()
		doSomething()
		doSomething()
		doSomething()
		time.Sleep(120 * time.Millisecond)
		fmt.Printf("with no goroutine : %d\n", time.Since(start1))
	*/
	/*
		wg := &sync.WaitGroup{}
		wg.Add(3)

		start2 := time.Now()
		go doSomething(wg)
		go doSomething(wg)
		go doSomething(wg)
		//time.Sleep(120 * time.Millisecond)
		wg.Wait() //Bottle neck
		fmt.Printf("with goroutine : %d\n", time.Since(start2))
	*/

	//infinity loop
	go func() {
		for {
			fmt.Println(read())
		}
	}()

	for i := 0; i < 10; i++ {
		write(i)
	}

}

func doSomething(wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("do something")
	wg.Done()
}

func write(n int) {
	mux.Lock()
	defer mux.Unlock()
	i = n

}

func read() int {
	mux.Lock()
	defer mux.Unlock()
	return i
}
