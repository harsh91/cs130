package main

import "fmt"

func main() {
	var largeNum int
	var smallNum int
	fmt.Print("Hey buddy tell me the largest number you are thinking of: ")
	fmt.Scan(&largeNum)
	fmt.Print("Now tell me the smallest number please: ")
	fmt.Scan(&smallNum)
	fmt.Print(largeNum%smallNum)
}
