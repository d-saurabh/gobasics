package main

import (
	"math"
	"sort"
)

func main()  {
	a:= []int{3,1,2,4,3}
	res:=Solution(a)
	println(res)
}

func Solution(A []int) int {
	// write your code in Go 1.4
	s := make([]int,0,5)

	var left int = 0
	for i:=0; i<len(A)-1;i++  {
		left+=A[i]
		var right int =0
		for j:=i+1;j<len(A) ;j++  {
			right+=A[j]
		}
		s = append(s,(int)(math.Abs(float64(left-right))))
		println((left-right))
	}
	sort.Ints(s)
	return s[0]
}