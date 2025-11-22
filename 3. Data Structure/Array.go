package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Alex"
	names[1] = "John"
	names[2] = "Mike"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// deklarasi langsung
	var values = [3]int{
		90, 80, 95,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(len(values[1:2]))
}
