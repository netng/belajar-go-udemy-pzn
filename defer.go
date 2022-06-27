package main

import "fmt"

func logging() {
	fmt.Println("logging")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("run application")
	result := 10 / value
	fmt.Println(result)
}

func main() {
	runApplication(0)
}
