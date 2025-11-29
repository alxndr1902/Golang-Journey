package main

import (
	"errors"
	"fmt"
)

func main() {
	hasil, err := Pembagian(10, 0)
	if err == nil {
		fmt.Println("HasiL: ", hasil)
	} else {
		fmt.Println("Error:", err.Error())
	}
}
func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi is zero")
	} else {
		return nilai / pembagi, nil
	}
}
