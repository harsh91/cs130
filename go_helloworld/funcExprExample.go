package main

import "fmt"

func main() {
	funcExp := func (value int) (int, bool) {
		return value/2, value%2 == 0
	}
	fmt.Println(funcExp(2))
}
