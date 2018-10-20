package main

import (
	"sync"
	"time"
)

var waitgroup sync.WaitGroup

func somefunc(s string) {
	for i := 0; i < 3; i++ {
		println(s)
		time.Sleep(time.Microsecond * 100)
	}
	waitgroup.Done()
}
func main() {
	waitgroup.Add(1)
	go somefunc("whats") // routine is a light weight thread
	waitgroup.Add(1)
	go somefunc("up") // another routine
	waitgroup.Wait()
}
