package main

func main()  {
	a:= []int{3,4,4,6,1,4,4}
	var o = Solution(5,a)
	for i:=range o  {
		println(o[i])
	}
}

func Solution(N int, A []int) []int {
	c:= make([]int,N,N)
	var max int = 0
	for i:= range A {
		if A[i]>=1 && A[i]<=N{
			c[A[i]-1]++
			x := c[A[i]-1]
			if x > max{
				max = x
			}
		} else if A[i] == N+1{
			for j:= range c  {
				c[j] =max
			}
		}
	}
	return c
}