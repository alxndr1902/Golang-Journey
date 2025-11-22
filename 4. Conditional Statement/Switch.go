package main

import "fmt"

func main() {
	name := "John"
	switch name {
	case "Alex":
		fmt.Println("Hello Alex")
	case "John":
		fmt.Println("Hello John")
	default:
		fmt.Println("You are not Alex or john")
	}

	// short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}
}
