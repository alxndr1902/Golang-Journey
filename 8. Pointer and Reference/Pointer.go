package main

import "fmt"

func main() {
	var address1 Address = Address{"Jakarta Barat", "Jakarta", "Indonesia"}
	var address2 *Address = &address1 // & = to point to address1 value, should address2 is altered, it will alter address1 value too
	address3 := &address1

	address2.City = "Jakarta Selatan"
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

}

type Address struct {
	City, Province, Country string
}
