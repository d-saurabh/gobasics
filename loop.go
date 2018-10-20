package main

func main()  {
	//onlyhave for loop

	//1.traditional for
	for i:=0;i<10 ;i++  {
		println(i)
	}

	//2. while
	j:=0
	for j<10 {

		println(j)
		j++
	}

	//3. infinite loop
	//for   {
	//	println("infinite loop")
	//}

	//4. kind of foreach
	var x = []int{1,2,3,4,5}
	for item := range x {
		println(x[item])
	}
	}