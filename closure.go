package main

import "fmt"

func main() {
	counter := 0
	name := "nandang"

	increment := func() {
		name := "kazu"
		fmt.Println("increment")
		counter += 1
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
