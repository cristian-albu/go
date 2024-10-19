package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	basics()
	arrays()
	maps()
	loops()
	slicesPerformance()
	stringsRunesAndBytes()
	structsAndInterfaces()
}

func basics() {
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

func arrays() {
	var intArr [3]int32 //  Allocates a contiguos memory of 4 byte size pieces, totaling in 12 bytes.
	intArr[1] = 123
	fmt.Println(intArr[0])   // index 0
	fmt.Println(intArr[1:3]) // 1 and 2
	fmt.Println(intArr)

	intArr2 := [3]int32{1, 2, 3}   // can also initialize it like this
	intArr3 := [...]int32{1, 2, 3} // or like this

	fmt.Println(&intArr2[0]) // the address of this memory, something like 0x000012120
	fmt.Println(intArr3)

	intSlice := []int32{4, 5, 6}                                                     // A slice is a wrapper around an array to make it more general and easier to use
	fmt.Printf("The length is %v with capacity %v \n", len(intSlice), cap(intSlice)) // len 3, cap 3
	fmt.Println(&intSlice[0])
	intSlice = append(intSlice, 7)                                                   // This will see if original array has enough space. If it does not it will create a new array with enough space and add the old values and the new value.
	fmt.Printf("The length is %v with capacity %v \n", len(intSlice), cap(intSlice)) // Capacity increased but you can't actually access the index of something outside the length. len 4 cap 6
	fmt.Println(&intSlice[0])                                                        // The memory address is now different from the initial one

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...) // Append multiple
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8) // Specify the length and the capacity
	fmt.Printf("The length is %v with capacity %v \n", len(intSlice3), cap(intSlice3))

	for i, v := range intSlice2 {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}
}

func maps() {
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

func loops() {
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

func timeloop(slice []int, n int) time.Duration {
	var t0 = time.Now()

	for len(slice) < n {
		slice = append(slice, 1)
	}

	return time.Since(t0)
}

func slicesPerformance() {
	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v \n", timeloop(testSlice, n)) // 16.58 ms
	fmt.Printf("Total time with preallocation: %v \n", timeloop(testSlice2, n))   // 4.16 ms
}

func stringsRunesAndBytes() {
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

type gasEngine struct {
	mpg     uint8
	gallons uint8
	// ownerInfo owner > this would be accessed as gasEngine.ownerInfo.name
	owner // this can be accessed as gasEngine.name
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type owner struct {
	name string
}

// assigning this functions to the gasEngine. This now becomes a "method". It has access to the fields and other methods of the struct
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

// This can me applied between others to be more generic.
type engine interface {
	milesLeft() uint8
}

// Now this can call the milesLeft method for different engine types
func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it")
	} else {
		fmt.Println("You can't make it")
	}
}

func structsAndInterfaces() {

	// you can ommit the names and values will be added in order
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15, owner: owner{"Alex"}}

	// you can also write it like this:
	myEngine.mpg = 35

	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)

	// anonymous struct
	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}

	fmt.Println(myEngine2)

	fmt.Println(myEngine.milesLeft())

	var myElectricEngine electricEngine = electricEngine{kwh: 35, mpkwh: 29}

	canMakeIt(myEngine, 25)
	canMakeIt(myElectricEngine, 35)
}
