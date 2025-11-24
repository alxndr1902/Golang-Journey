package main

import "fmt"

func main() {
	goodBye := getGoodBye
	fmt.Print(goodBye("alex"))
}

func getGoodBye(name string) string {
	return "Good bye " + name
}
