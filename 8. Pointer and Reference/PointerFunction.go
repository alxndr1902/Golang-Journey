package main

import "fmt"

func main() {
	address := Address{}
	ChangeCountryToIndonesia(&address)

	// you can be like this
	// var address *Address = &Address{}
	// ChangeCountryToIndonesia(address)
	fmt.Println(address)
}

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}
