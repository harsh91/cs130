package main

import "fmt"


type candyStore struct {
	Name string
	Candies []string
}

type candyFullStore struct {
	Name string
	Candies map[string]int
}

func main() {
	struct1 := candyStore {
		Name: "Harsh",
		Candies: []string{"Snickers", "KitKat", "Cadbury"},
	}
	fmt.Println(struct1)

	 struct2 := candyFullStore{
		 Name: "Harsh",
		 Candies: map[string]string {
			 "Kitkat": "20",
			 "Snickers": "10",
			 "Cadbury": "23",
		 },
	 }
	fmt.Println(struct2)
}