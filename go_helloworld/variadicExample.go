package main

import "fmt"

func main() {
	biggestNumber := variadicGreatest(2,10,2,3,4,6,92,1,43,32)
	fmt.Println(biggestNumber)
}

func variadicGreatest(numbers ...int) int {
	var largest int
	for _,num := range numbers {
		if num > largest {
			largest = num
		}
	}
	return largest
}