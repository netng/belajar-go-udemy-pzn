package main

import "fmt"

func getFullName() (string, string) {
	return "nandang", "sopyan"
}

func main() {
	// _ mean, the return value for its will ignored
	_, lastName := getFullName()
	fmt.Println(lastName)
}
