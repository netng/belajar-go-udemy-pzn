package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan nol")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	result, err := Pembagian(10, 2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
