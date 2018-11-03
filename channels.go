package main

func some(c chan int,value int){
	c <- value*2
}
func main()  {
	//use channel with go routines
	//channel := make(chan int)
	//
	//go some(channel,4)
	//go some(channel,10)

	//val1 := <- channel
	//val2 := <- channel
	//or
	//val1,val2 := <-channel,<-channel
	//fmt.Println(val1,val2)

	results := make(chan int,10)
	jobs := make(chan int, 10)

	go worker(jobs,results)

	for i:= 2; i<7;i++  {
		jobs <- i
	}

	close(jobs)

	for i:= 0; i<5;i++  {
		println(<-results)
	}
}

func worker(jobs <-chan int,results chan<- int){
	for n:= range jobs  {
		results <- fib(n)
	}
}

func fib(no int)int  {
	if no<=1{
		return no
	}

	return fib(no-1)+fib(no-2)
}