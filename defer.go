package main

import "fmt"

func log() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer log()
	result := 10 / value
	fmt.Println("Application running...")
	fmt.Println(result)

}

func main() {
	runApplication(0)
}
