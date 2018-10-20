package main

func main() {
	//map is a reference type
	//var scores map[string]float32 //[key]valuetype
	//or
	scores := make(map[string]int) //make initialize the map
	scores["saurabh"] = 70
	scores["lucas"] = 79
	scores["dante"] = 77

	loop(scores)
	delete(scores, "saurabh")
	loop(scores)

}

func loop(c map[string]int) {
	for key, value := range c {
		println(key, value)
	}
}
