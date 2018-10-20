package main

import (
	"sync"
	"time"
)

func foo() {
	//defer follows LIFO approach
	defer println("first") // only get executed when rest of the foo is done
	defer println("second")
	println("foo is executing")
}

var waitgroup sync.WaitGroup

func somefunc(s string) {
	defer waitgroup.Done()
	for i := 0; i < 3; i++ {
		println(s)
		time.Sleep(time.Microsecond * 100)
	}

}
func main() {
	foo()
	waitgroup.Add(1)
	go somefunc("whats") // routine is a light weight thread
	waitgroup.Add(1)
	go somefunc("up") // another routine
	waitgroup.Wait()
}
