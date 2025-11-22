package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "alex"
	// person["address"] = ""

	person := map[string]string{
		"name":    "alex",
		"address": "earth",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	delete(person, "name")
	fmt.Println(person)

}
