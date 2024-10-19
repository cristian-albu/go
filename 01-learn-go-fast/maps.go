package main

import "fmt"

func Maps() {
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45, "John": 64}
	fmt.Println(myMap2["Adam"])  // 23
	fmt.Println(myMap2["Jason"]) // 0, default type if it does not exist

	var age, ok = myMap2["Jason"] // 0, false - second value is if the key exists
	fmt.Println(age, ok)

	fmt.Println(myMap2["Sarah"]) // 45
	delete(myMap2, "Sarah")      // Delete by reference. Does not return anything
	fmt.Println(myMap2["Sarah"]) // 0

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, age: %v\n", name, age) // looping over map. No order is preserved
	}
}
