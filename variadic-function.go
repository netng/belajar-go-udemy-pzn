package main

import "fmt"

func sumAll(numbers ...int) (total int) {
	total = 0
	for _, number := range numbers {
		total += number
	}
	return
}

func main() {
	total := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(total)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	total = sumAll(slice...)
	fmt.Println(total)
}
