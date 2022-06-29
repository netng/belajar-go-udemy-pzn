package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hi", customer.Name, ", my name is", name)
}

func main() {
	customer1 := Customer{
		Name:    "kazu",
		Address: "bogor",
		Age:     20,
	}

	customer1.sayHi("nandang")

	// fmt.Println(customer1)
	// fmt.Println(customer1.Name)
	// fmt.Println(customer1.Address)
	// fmt.Println(customer1.Age)

	// customer1.Age = 30
	// fmt.Println(customer1.Age)
}
