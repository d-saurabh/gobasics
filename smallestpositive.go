package main

import "sort"

func main()  {
	a:= []int{-3,-2,1}
	//a:= []int{1,-3,-2,}
	//a:= []int{-1,-3}
	//a:= []int{1,3,6,4,1,2,3}
	//a:= []int{1}
	//a:= []int{-1}
	//a:= []int{2}
	o := Solution(a)
	println(o)
}

func Solution(A []int) int {

	if(len(A)==1){
		if A[0]<=0{
			return 1
		} else if A[0] > 1 {
			return A[0] - 1
		}
	}

	min:=1
	sort.Ints(A)

	for i:= range A{
		if i < len(A)-1{
			j:= A[i+1]-A[i]
			if j>1{
				min = A[i+1]-1
			}
		} else{
			if min <0{
				min = 1
			} else if min >=0 && min <= 1{
				min = A[i]+1
			} else if min ==0{
				min =1
			}

		}
	}
	return min
}