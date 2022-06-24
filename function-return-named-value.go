package main

import "fmt"

func getFullName2() (firstName, lastName string) {
	firstName = "nandang"
	lastName = "sopyan"

	return
}

func main() {
	firstName, lastName := getFullName2()
	fmt.Println(firstName)
	fmt.Println(lastName)
}
