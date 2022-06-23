package main

import "fmt"

func main() {
	name := "kazunius"

	if name == "nandang" {
		fmt.Println("Hi, nandang")
	} else if name == "kazu" {
		fmt.Println("Hi, kazu")
	} else {
		fmt.Println("Hi, kenalan dong")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
