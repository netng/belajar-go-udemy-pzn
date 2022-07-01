package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHi(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Users struct {
	Name string
}

func (u Users) GetName() string {
	return u.Name
}

func main() {
	user := Users{
		Name: "nandang",
	}

	sayHi(user)
}
