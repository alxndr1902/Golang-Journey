package main

import "fmt"

func main() {
	sayHello()
	sayHelloTo("Alex", "S")
	result := getHello("alex")
	fmt.Println(result)

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	namaDepan, _ := getFullName()
	fmt.Println(namaDepan)

	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}

func sayHello() {
	fmt.Println("Hello")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func getHello(name string) string {
	return "Hello " + name
}

func getFullName() (string, string) {
	return "Alex", "S"
}

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "John"
	middleName = "F."
	lastName = "Kennedy"

	return firstName, middleName, lastName
}
