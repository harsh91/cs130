package main

import "fmt"

type person struct {
	name string
	age int
	email string
}

type pets struct {
	name string
	pet []string
}

func (p person) String() string {
	return fmt.Sprintf("Name : %s, Age : %d, Email : %s", p.name, p.age, p.email)
}

func main() {
	fmt.Println("Struct Example")
	var x person
	x.name = "Harsh"
	x.age = 24
	fmt.Println(x)

	y := person {
		"Deepti",
		23,
		"deepti@gmail.com",
	}
	fmt.Println("Printing structs as composite literal: ", y)

	z := person {
		name : "Rahul",
		age : 24,
	}
	fmt.Println("Printing struct with less arguements: ", z)
	z.email = "rahul@gmail.com"
	fmt.Println("After adding a new value to struct: ", z)

	fmt.Println(z)

	a := pets {
		"Harsh",
		[]string{"dog", "cat"},
	}
	fmt.Println("Struct with a slice inside it in it: ", a)
}
