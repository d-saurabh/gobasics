package main

import "fmt"

func some(c chan int,value int){
	c <- value*2
}
func main()  {
	//use channel with go routines
	channel := make(chan int)

	go some(channel,4)
	go some(channel,10)

	//val1 := <- channel
	//val2 := <- channel
	//or
	val1,val2 := <-channel,<-channel
	fmt.Println(val1,val2)
}
