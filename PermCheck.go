package main

import "sort"
import "github.com/golang/example/stringutil"

func main()  {
	s := stringutil.Reverse("saurabh")
	println(s)
	a:= []int{1,2,3}
	b := a[:0]
	for _, x := range a {
		if x<10 {
			b = append(b, x)
		}
	}

	x:= a[0:]
	for i:= range x{
		println(i)
	}
	var o = Solution(a)
	println(o)
}

func Solution(A []int) int {
	sort.Ints(A)
	for i:=0;i<len(A) ;i++  {
		if i != (len(A)-1) &&  A[i+1]-A[i] != 1{
			return 0
		}
	}
	return 1
}