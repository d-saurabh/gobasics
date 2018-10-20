package main

type person struct {
	name   string
	age    int
	gender string
}

//value receivers
func (p person) valuefunc() string {
	return p.name
}

//pointer receiver
func (p *person) pointerfunc() {
	p.name = p.name + "deshpande"
}
func main() {
	a_person := person{name: "saurabh",
		age:    30,
		gender: "male"}

	println(a_person.valuefunc())
	a_person.pointerfunc()
	println(a_person.name)
}
