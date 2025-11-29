package main

import "fmt"

func main() {
	alex := Man{"Alex"}
	alex.Married()
	fmt.Println(alex.Name)

}

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}
