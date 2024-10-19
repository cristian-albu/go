package main

import (
	"fmt"
)

func Loops() {
	var i int = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i = i + 1
	}

	j := 0
	for j < 10 {
		fmt.Println(j)
		j += 1
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
