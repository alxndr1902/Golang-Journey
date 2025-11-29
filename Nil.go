package main

import "fmt"

func main() {
	data := NewMap("alex")
	if data == nil {
		fmt.Println("Data map is still empty")
	} else {
		fmt.Println(data["name"])
	}
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
