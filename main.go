package main

import "fmt"

func main() {
	add := AddNumber(10, 5)
	min := MinNumber(10, 5)

	fmt.Println(add)
	fmt.Println(min)
	fmt.Println("SUCCESS 2")
}

func AddNumber(a, b int) int {

	return a + b
}

func MinNumber(a, b int) int {

	return a - b
}
