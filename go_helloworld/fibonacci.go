package main

import "fmt"

func main() {
	fmt.Print(fibonacci(5))
}

func fibonacci(number int) int {
	fibonacci := number
	for i := number-1 ; i > 0 ; i-- {
		fibonacci = fibonacci * i
	}
	return fibonacci;
}