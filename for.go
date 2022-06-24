package main

import (
	"fmt"
)

func main() {
	//var count int8 = 1

	//for count <= 10 {
	//	fmt.Println("Perulangan ke", count)
	//	count++
	//}

	//for count := 1; count <= 10; count++ {
	//	fmt.Println("perulangan ke", count)
	//}

	slice := []string{
		"ujang",
		"udin",
	}

	for i, name := range slice {
		fmt.Println("Index ke ", i, " = ", name)
	}

	maps := map[string]string{
		"name":  "ujang",
		"title": "programmer",
	}

	for key, name := range maps {
		fmt.Println(key, ": ", name)
	}
}
