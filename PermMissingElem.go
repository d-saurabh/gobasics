package main

import (
	"sort"
)

func main()  {
	a:= []int{1,2,3,5}
	var o = Solution(a)
	println(o)
}

func Solution(A []int) int{
	sort.Ints(A)
	for i:= range A {
		if i != len(A) && A[i+1]-A[i] != 1 {
			return A[i] + 1
		} else if i == len(A) {
			return A[i] - 1
		}
	}
	return 0
}

