package main

type person struct {
	name   string
	age    int
	gender string
}

func main() {
	a_person := person{name: "saurabh",
		age:    30,
		gender: "male"}

	println(a_person.age)
	println(a_person.gender)
	println(a_person.name)
}
