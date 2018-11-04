package main

func main()  {
	a:= []int{1,3,1,4,2,3,5,4}

	var o = Solution(6,a)
	println(o)
}

func Solution(X int, A []int) int {
	for i:= range A  {
		if A[i] == X{
			return i
		}
	}
	return -1
}
