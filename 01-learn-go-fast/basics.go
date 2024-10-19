package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func printMe(printValue string) {
	fmt.Println(printValue)
}

func multipleReturns() (int, int) {
	return 1, 2
}

func intDivision(numverator int, denuminator int) (int, error) {
	var err error

	if denuminator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, err
	}
	return numverator / denuminator, err
}

func Basics() {
	var intNum int = 54
	fmt.Println(intNum)

	var floatNum float64 = 32.5
	fmt.Println((floatNum))

	var result float64 = floatNum + float64(intNum)
	fmt.Println(result)

	var myString string = "Hello" + " " + "World"
	var myString2 string = "Hello \n world"
	var myString3 string = `Hello
	World`

	fmt.Println(myString, myString2, myString3)
	fmt.Println(len("A"))                        // 1 - the legth of bytes
	fmt.Println(utf8.RuneCountInString(("len"))) // 3

	var myRune rune = 'a' // 97 > ASCII
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	myVar := "text" // infered
	fmt.Println(myVar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2) // declare multiple inline

	var1 = var1 + 5
	fmt.Println(var1)

	const myConst int = 1 // can't change it

	fmt.Println(myConst)

	first, second := multipleReturns()

	printMe("i was printed")
	divisionResult, err := intDivision(4, 2)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("The division %v and the other two %v, %v", divisionResult, first, second)

	divisionResult, err = intDivision(10, 2) // overwrite the previous return values
	switch {
	case err != nil:
		fmt.Println(err)
	case divisionResult == 5:
		fmt.Println(true)
	default:
		fmt.Println("Can't reach this")
	}

	switch divisionResult {
	case 5:
		fmt.Println("This is a conditional switch")
	}

}
