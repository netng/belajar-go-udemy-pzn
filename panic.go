package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error with message: ", message)
	}
	fmt.Println("Finish")
}

func runApp(err bool) {
	defer endApp()
	if err {
		panic("APPLICATION ERROR")
	}
	fmt.Println("Application running")
}

func main() {
	runApp(true)
	fmt.Println("nandang")
}
