package main

import "fmt"

func main() {
	name := "alexxsss"

	if name == "alex" {
		fmt.Println("Hello", name)
	} else if name == "john" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("you are not alex or john")
	}

	//short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
