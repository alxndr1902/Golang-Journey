package main

import "fmt"

func main() {
	fmt.Println(sumAll(10, 10, 10))
	fmt.Print(sumAllSlice([]int{10, 10, 10}), "\n")

	numbers := []int{10, 10, 10}
	fmt.Println(sumAll(numbers...))
}

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func sumAllSlice(numbers []int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}
