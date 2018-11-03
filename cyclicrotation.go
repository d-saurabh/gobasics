package main

func main()  {
	a:= []int{1,2,3,4}
	var o = Solution(a,3)
	for index:= range o {
		println(a[index])
	}
}

func Solution(A []int, K int) []int{
	if len(A) <=1{
		return A
	}

	l := len(A)
	temp := make([]int,l,l)

	for K>0{
		var j = 0
		temp[j] = A[l-1]
		j++
		for i:=0;i<=l-2 ;i++  {
			temp[j] = A[i]
			j++
		}
		copy(A, temp)
		K--
	}
	return temp

}