package main

import "fmt"

func main() {
	// Break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 8 {
			break
		}

	}

	// Continue
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
	}
}
