package main

func main()  {

	a :=  []int{9, 3, 9, 3, 9, 7, 9}
	var o = Solution(a)
	println(o)
}

func Solution(A []int) int {
	s := make([]int, len(A),len(A))
	var unPaired = -1
	for i:=0;i<len(A) ;i++  {
		var isUnpaired bool = true
		for j:=i+1;j<len(A) ;j++  {
			if A[i] == A[j]{
				isUnpaired =false
				s = append(s,A[i])
				break
			}
		}
		if isUnpaired && !Contains(s,A[i]) {
			unPaired = A[i]
			break
		}
	}
	return unPaired
}

func Contains(A []int,item int)bool  {
	for index := range A{
		if A[index] == item{
			return true
		}
	}
	return false
}