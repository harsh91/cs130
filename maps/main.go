package main

import "fmt"

func main() {
	scoreMap := map[string]string {
		"Harsh": "100",
		"Abhas": "90",
		"Deepti": "80",
		"Sarika": "95",
		"Jeev": "91",
	}
	fmt.Println("Maps in intially looks like")
	fmt.Println(scoreMap)
	scoreMap["Mukasir"] = "70"
	fmt.Println("Maps after adding an element looks like")
	fmt.Println(scoreMap)
	fmt.Print("Deleting : ")
	fmt.Println(scoreMap["Mukasir"])
	delete(scoreMap, "Mukasir")
	fmt.Println("After deleting map looks like")
	fmt.Println(scoreMap)
	fmt.Println("Modfying the existing value")
	scoreMap["Harsh"] = "99"
	for keys, values := range scoreMap {
		fmt.Println(keys + ": " + values)
	}
}
