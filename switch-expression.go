package main

import "fmt"

func main() {
	name := "nan"

	switch name {
	case "nandang":
		fmt.Println("Hi nandang")
		fmt.Println("Hi nandang")
	case "kazu":
		fmt.Println("Hi kazu")
	default:
		fmt.Println("Kenalan yuk")
	}

	//switch length := len(name); length > 5 {
	//case true:
	//	fmt.Println("Nama terlalu panjang")
	//case false:
	//	fmt.Println("Nama sudah benar")
	//}

	length := len(name)

	switch {
	case length > 5:
		fmt.Println("Nama terlalu panjang")
	case length <= 5 && length >= 3:
		fmt.Println("Nama sudah benar")
	default:
		fmt.Println("Nama salah")
	}
}
