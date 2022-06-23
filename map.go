package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "nandang",
		"city": "bogor",
	}

	fmt.Println(person)

	person["job"] = "programmer"
	fmt.Println(person)

	delete(person, "job")
	fmt.Println(person)

	fmt.Println(len(person))

	fmt.Println(person["name"])
	fmt.Println(person["city"])

}
