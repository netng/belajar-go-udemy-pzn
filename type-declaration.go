package main

import "fmt"

func main() {
	type (
		NoKtp   string
		Married bool
	)

	var noKtp NoKtp = "1234567890"
	var marriedStatus Married = true

	fmt.Println(noKtp)
	fmt.Println(marriedStatus)
}
