package main

import "fmt"

func main() {
	const firstName string = "ujang"
	const lastName = "go"
	const value = 100

	//firstName = "ucup" //<-- akan error karena const tidak bisa di ubah value nya
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	const (
		language string = "indonesia"
		version         = 1
	)

	fmt.Println(language)
	fmt.Println(version)

}
