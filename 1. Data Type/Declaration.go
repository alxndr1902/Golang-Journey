package main

import "fmt"

func main() {
	type NoKTP string

	var KTPAlex NoKTP = "123123"
	fmt.Println(KTPAlex)
	fmt.Println(NoKTP("no ktp"))
}
