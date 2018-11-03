package main

import "math"

func main()  {

	var o = Solution(10,85,30)
	println(o)
}
func Solution(X int, Y int, D int) int {
	if X >= Y{
		return 0
	}

	res := (float64)(Y - X)/float64(D)

	x:= (int)(math.Ceil(res))

	return x

}