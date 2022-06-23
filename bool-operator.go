package main

import "fmt"

func main() {
	nilaiUjian := 90
	nilaiAbsensi := 80

	lulusUjian := nilaiUjian > 80
	lulusAbsensi := nilaiAbsensi > 80
	lulus := lulusUjian && lulusAbsensi

	fmt.Println(lulus)

}
