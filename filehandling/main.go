package main

import (
	"fmt"
	"encoding/json"
	"os"
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

	fileName := "new.json"
	_, fileErr := os.Stat(fileName)
	var handle, _ *os.File
	if fileErr == nil {
		handle, _ = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
		fmt.Println("Opening")
	} else {
		handle, _ = os.Create(fileName)
	}
	handle.Write([]byte(jsonData))
	fmt.Println("Written data to file!!")
}