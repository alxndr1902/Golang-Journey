package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("Perulangan ke ", i)
	// }

	names := []string{"alex", "john", "doe"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
}
