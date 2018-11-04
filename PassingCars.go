package main

func main(){
	a:= []int{0,1,0,1,1,0,1}
	o := Solution(a)
	println(o)
}

func Solution(A []int) int {
	if len(A)==1{
		return 0
	}
	count:=0
	for i:= range A{
		if A[i] ==0 && i<len(A)-1{

			for j:= i+1;j<len(A);j++{
				if A[j] == 1{
					count++
					if count > 1000000000{
						return -1
					}
				}
			}
		}
	}
	return count
}