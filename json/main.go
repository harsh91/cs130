package main

import (
	"fmt"
	"encoding/json"
)


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
		Candies: map[string]int {
			"Kitkat": 20,
			"Snickers": 10,
			"Cadbury": 23,
		},
	}
	fmt.Println(struct2)

	jsonData, _ := json.Marshal(struct2)
	fmt.Println("Json Data: ", string(jsonData))

	var candy candyFullStore
	err := json.Unmarshal([]byte(jsonData), &candy)
	if err == nil {
		fmt.Println("Decoded json data : ", candy)
	}

}