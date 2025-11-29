package main

import "fmt"

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}

func Ups() any {
	return 1
}
