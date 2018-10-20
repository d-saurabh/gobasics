package main

type T struct {
	name string // name of the object
	value int // its value
}

//channels
func foo(c chan int, someValue int) {
	c <- someValue*5
}

func add(x float64, y float64) float64{
	return x +y
}

//pointers

func main(){

	//pointer & memory address and * read from the address
	x := 15
	a := &x
	println(a)
	println(*a)
	*a  = 5
	println(x)
	*a = *a**a
	println(x)
	//var num1 float64= 5.6
	//var num2 float64 = 9.5
	//
	//println(add(num1,num2))

	//channel
	//fooVal := make(chan int)
	//go foo(fooVal ,5)
	//go foo(fooVal,6)
	//
	//v1 := <- fooVal
	//v2:= <- fooVal
	//
	//println(v1)
	//println(v2)
}
/*
func main(){
	fmt.Println("hello world")
	var strs []string

	strs = append(strs,"hello")
	strs = append(strs,"world")

	for _, item := range strs  {

		fmt.Println(item)
	}

	//simple if
	if 1 >0 {
		fmt.Println("small")
	}

	//if with extra initialization param
	if err := "error"; err == "error"{
		fmt.Println("its an error")
	}

	//every function returns an err as  a return in go
	// alsways check err !=nil as Go doesent have exception handling
	/*f, err := os.Open(name)
	if err != nil {
		return err
	}
	d, err := f.Stat()
	if err != nil {
		f.Close()
		return err
	}
	codeUsing(f, d) */

	//Redeclaration and reassignment
	// f, err := os.Open(name)  Note: err is declared here
	//d, err := f.Stat() Note: err is reassigned and not declared here as it is created in above statement and is in the same scope

	//For
	//1. Like a C for
	// for init; condition; post { }
	//2. Like a C while
	// for condition { }
	//3. Like a C for(;;)
	// for { }
	//sum := 0
	//for i:=0;i<10 ;i++  {
	//	sum+=i
	//}
	//fmt.Println(sum)

	//loop over array, slice, string or map or reading from a channel
	//for key, value := range oldMap {
	//		newMap[ket] = value
	//}

	//If you only need the first item in the range (the key or index), drop the second:
	//for key := range m {
	//	if key.expired(){
	//		delete(m,key)
	//	}
	//}

	//If you only need the second item in the range (the value), use the blank identifier, an underscore, to discard the first
	//sum := 0
	//for _, value := range array {
	//	sum += value
	//}


	//For strings, the range does more work for you, breaking out individual Unicode code points by parsing the UTF-8. Erroneous encodings consume one byte and produce the replacement rune U+FFFD. (The name (with associated builtin type) rune is Go terminology for a single Unicode code point. See the language specification for details.) The loop
	//for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
	//	fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	//}

	// Reverse a
	//for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
	//	a[i], a[j] = a[j], a[i]
	//}


	//Switch


	//Multiple return values ¶

	//Defer
	//Go's defer statement schedules a function call (the deferred function) to be run immediately before the function executing the defer returns. It's an unusual but effective way to deal with situations such as resources that must be released regardless of which path a function takes to return. The canonical examples are unlocking a mutex or closing a file.
	//Deferred functions are executed in LIFO order,

	//Data
	//Allocation with new
	//Go has two allocation primitives, the built-in functions new and make.
	//NEW:
	//It's a built-in function that allocates memory, but unlike its namesakes in some other languages it does not initialize the memory, it only zeros it
	//That is, new(T) allocates zeroed storage for a new item of type T and returns its address, a value of type *T.
	//In Go terminology, it returns a pointer to a newly allocated zero value of type T.

	//p := new(SyncedBuffer)  // type *SyncedBuffer
	//var v SyncedBuffer      // type  SyncedBuffe

	//Constructors and composite literals

	//Allocation with make
	//Back to allocation. The built-in function make(T, args) serves a purpose different from new(T)
	//t creates slices, maps, and channels only, and it returns an initialized (not zeroed) value of type T (not *T)
	//he reason for the distinction is that these three types represent, under the covers, references to data structures that must be initialized before use
	//A slice, for example, is a three-item descriptor containing a pointer to the data (inside an array), the length, and the capacity, and until those items are initialized, the slice is nil
	//For slices, maps, and channels, make initializes the internal data structure and prepares the value for use. For instance,
	//ints := make([]int, 10, 100)

	//Arrays

	//Slices
	//var x =[]int{1,2}
	//var y = [][]int

	//Maps (dictionary)
	/*var timeZone = map[string]int{
		"UTC":  0*60*60,
		"EST": -5*60*60,
		"CST": -6*60*60,
		"MST": -7*60*60,
		"PST": -8*60*60,
	}
	offset := timeZone["EST"]

	//Append
	x := []int{1,2,3}
	y := []int{4,5,6}
	x = append(x, y...)
	fmt.Println(x)
*/



//}*/