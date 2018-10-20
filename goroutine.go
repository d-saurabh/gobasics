package main

import "time"

func somefunc(s string) {
	for i := 0; i < 3; i++ {
		println(s)
		time.Sleep(time.Microsecond * 100)
	}
}
func main() {
	go somefunc("whats") // routine is a light weight thread
	somefunc("up")       // if this is also set to go routine then nothing will be printed as, your program finished before your routine start

}
