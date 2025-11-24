package main

import "fmt"

func main() {
	runApp(true)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Something went wrong")
	}
}

func endApp() {
	fmt.Println("End app")
	message := recover()
	fmt.Println("Recovered message:", message)
}
