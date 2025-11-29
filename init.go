package main

import (
	"fmt"
	"golang-tutorial/database"
	_ "golang-tutorial/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
