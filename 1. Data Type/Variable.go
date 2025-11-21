package main

import "fmt"

func main() {
	var name string
	name = "Alex"
	fmt.Println(name)

	name = "Alexander"
	fmt.Println(name)

	gender := "Laki-laki"
	fmt.Println(gender)

	//Multiple variable
	var (
		firstName = "Alex"
		lastName  = "S"
	)
	fmt.Println("Hallo my name is", firstName, lastName)
}
