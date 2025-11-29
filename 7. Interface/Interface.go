package main

import (
	"fmt"
)

func main() {
	person := Person{Name: "Alex"}
	sayHello(person)
	hello := person.sayHello()
	fmt.Println(hello)
}

type Person struct {
	Name string
}

type hasName interface {
	getName() string
	sayHello() string
}

func sayHello(value hasName) {
	fmt.Println("hello", value.getName())
}

func (person Person) getName() string {
	return person.Name
}

func (person Person) sayHello() string {
	return "Hello, my name is " + person.Name
}
