package main

import "fmt"

func sayGo(value int) interface{} {
	if value == 3 {
		return true
	} else {
		return "Halo"
	}
}

func main() {
	saygo := sayGo(2)
	saygo2 := sayGo(3)
	fmt.Println(saygo)
	fmt.Println(saygo2)
}
