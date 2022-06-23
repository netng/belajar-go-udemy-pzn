package main

import "fmt"

func main() {
	months := [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	fmt.Println(months)

	slice1 := months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	slice1[0] = "mei aja"

	slice2 := months[4:]

	slice3 := append(slice2, "ujang")
	fmt.Println(slice3)
	slice3[0] = "mei ok"

	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	anotherSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(anotherSlice)

	justAnotherSlice := make([]string, 2, 2)
	fmt.Println(justAnotherSlice)
	fmt.Println(len(justAnotherSlice))
	fmt.Println(cap(justAnotherSlice))

	justAnotherSlice[0] = "udin"
	justAnotherSlice[1] = "ucup"
	justAnotherSlice2 := append(justAnotherSlice, "aep")
	fmt.Println(justAnotherSlice2)
	fmt.Println(len(justAnotherSlice2))

	justAnotherSlice2[0] = "udin penyok"
	fmt.Println(justAnotherSlice2)
	fmt.Println(justAnotherSlice)

}
