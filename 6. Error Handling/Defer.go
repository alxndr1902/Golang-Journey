package main

import "fmt"

func main() {
	runApplication()
}

func runApplication() {
	defer logging()
	fmt.Println("Run application")
}

func logging() {
	fmt.Println("Selesai memanggil function")
}
