package main

import (
	"fmt"
)

func main() {

	var input string
	fmt.Print("Want to run basics? y/n: ")
	fmt.Scan(&input)

	if input == "y" {
		Basics()
		Arrays()
		Maps()
		Loops()
		StringsRunesAndBytes()
		StructsAndInterfaces()
		Pointers()
		Goroutines()
		Channels()
		Generics()
	} else {
		fmt.Print("Okay then")

	}

}
