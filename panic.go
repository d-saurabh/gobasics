package main

import (
	"sync"
	"time"
)

//panic: program will stop running
//recover: recover from panic situation

var waitgroup sync.WaitGroup

func cleanup() {
	defer waitgroup.Done()
	if r := recover(); r != nil {
		println("cleanup executing", r)
	}

}

func somefunc(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {

		time.Sleep(time.Microsecond * 100)
		println(s)

		if i == 2 {
			panic("ohh boy its 2")
		}
	}

}
func main() {
	waitgroup.Add(1)
	go somefunc("whats") // routine is a light weight thread
	waitgroup.Add(1)
	go somefunc("up") // another routine
	waitgroup.Wait()
}
