package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "nan"
	names[1] = "dang"
	names[2] = "com"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	values := [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	var anotherArray [10]string
	fmt.Println(len(anotherArray))
}
