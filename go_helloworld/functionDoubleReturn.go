package main

import "fmt"

func main() {
	var halfValue int
	var evenOrNot bool
	halfValue, evenOrNot = half(2)
	fmt.Println(halfValue,evenOrNot)
}

func half(value int) (int, bool) {
	return value/2, value%2 == 0
}
