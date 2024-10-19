package main

import (
	"fmt"
	"strings"
)

func StringsRunesAndBytes() {
	var myString = "Résumé"
	var indexed = myString[1]
	var indexed2 = myString[2]
	fmt.Printf("Indexed: %v %T\n", indexed, indexed)   // 195. It is Actually 233 but it does not fit into uint8 (0-255) because the encoding needs some bits for the padding;
	fmt.Printf("Indexed: %v %T\n", indexed2, indexed2) // 169. This is the rest of the information about the same character "é"

	for i, v := range myString {
		fmt.Println(i, v) //0 82 1 233 Notice there is no 2 cause it takes two uint8 for the character 3 115 4 117 5 109 6 233 > Last index now is 6
	}

	// easier to use runes
	var myString2 = []rune("Résumé")
	var indexed3 = myString[1] // 82
	fmt.Printf("Indexed: %v %T\n", indexed3, indexed3)
	for i, v := range myString2 {
		fmt.Println(i, v) //0 82 1 233 2 115 3 117 4 109 5 233 > Last index now is 5
	}

	var myRune = 'a'    // single quotes
	fmt.Println(myRune) // 97 (ASCII)

	var strSlice = []string{"s", "u", "b"}

	var catStr = ""
	fmt.Println(catStr)
	// catStr[0] = "a" can't assign once it's created

	for i := range strSlice {
		catStr += strSlice[i]
	}

	// but when you range of one it created a new string and overwrite the whole variable and not the indexed values. It makes a new string in each loop. Not efficient
	fmt.Println(catStr)

	// This allocates a new array and build the whole thing and then assigns it. Much faster
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}

	var catStr2 = strBuilder.String()
	fmt.Printf("%v\n", catStr2)
}
