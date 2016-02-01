package main

import "fmt"

func main() {
	var userName string
	fmt.Print("Please enter user name: \n");
	fmt.Scan(&userName);
	fmt.Println("Hello " + userName);
}