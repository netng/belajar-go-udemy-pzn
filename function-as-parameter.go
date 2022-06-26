package main

import (
	"fmt"
	"strings"
)

type Filter func(string) string
type ToUpperCase func(string) string

func sayHelloWithFilter(name string, filter Filter, toUpper ToUpperCase) {
	nameFiltered := toUpper(filter(name))
	fmt.Println("Hello ", nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func toUpperCase(name string) string {
	return strings.ToUpper(name)
}

func main() {
	sayHelloWithFilter("nandang", spamFilter, toUpperCase)
	sayHelloWithFilter("anjing", spamFilter, toUpperCase)
}
