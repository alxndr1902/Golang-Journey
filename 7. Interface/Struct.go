package main

import "fmt"

func main() {
	var alex Customer
	alex.Name = "Alex"
	alex.Address = "earth 123"
	alex.Age = 23

	fmt.Println(alex)
	fmt.Println(alex.Name)
	fmt.Println(alex.Address)
	fmt.Println(alex.Age)

	JohnDoe := Customer{
		Name:    "John Doe",
		Age:     42,
		Address: "Mars",
	}
	fmt.Println(JohnDoe)
	JohnDoe.sayHello("budi")
}

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}
