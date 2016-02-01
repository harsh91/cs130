package main

import (
	"fmt"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age	int
	Party bool
}

func main() {
	p1 := person{
		Name : "Harsh",
		Age : 24,
	}
	if p1.Age > 21 {
		p1.Party = true
	}
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
	}
	err = tpl.Execute(os.Stdout, p1)
	if err != nil {
		fmt.Println(err)
	}
}