package main

import "fmt"

func Pointers() {
	var p *int32 = new(int32) // this holds the memory address to a int32 value
	var i int32

	fmt.Printf("The value p points to is: %v", *p) // dereferencing the pointer *p > 0 . The memory location is zeroed out when it has nothing assigned to it
	fmt.Printf("\nThe value if i is: %v\n", i)     // 0

	*p = 10 // this sets the value at that memory location to 10
	// If new(int32) is not used the pointer will be nil and trying to assign a value to it will give a nil pointer exception

	p = &i // Now p hold the memory address of i
	*p = 1 // Now the value of i is 1

	var slice = []int32{1, 2, 3}
	var sliceCopy = slice

	sliceCopy[2] = 4

	fmt.Println(slice)     // 1 2 4
	fmt.Println(sliceCopy) // 1 2 4

	// That as a copy by reference when only the pointers changes but the underlying array values are the same so when you change on you change the other

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The memory location of the thing1 array is: %p\n", &thing1) // 0xc00001e270
	var result [5]float64 = square(thing1)
	fmt.Printf("The result is: %v\n", result)

	var thing3 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The memory location of the thing3 array is: %p\n", &thing3) // 0xc00001e300
	var result2 [5]float64 = squareWithPointer(&thing3)
	fmt.Printf("The result2 is: %v\n", result2)

	// So the easy way to remember is:
	// "*" is the Dereference Operator > *ptr gives you the value stored at the address ptr is pointing to.
	// "&" Address-of Operator > &x gives you the memory address of x, turning x into a pointer to that address.

}

// We are actually working with 2 different arrays. A copy is made when the function is called
func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("The memory location of the thing2 array is: %p\n", &thing2) // 0xc00001e2a0 > different.
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}

	return thing2
}

// Array by reference. No copy is made
func squareWithPointer(thing4 *[5]float64) [5]float64 {
	fmt.Printf("The memory location of the thing4 array is: %p\n", thing4) // 0xc00001e300
	for i := range thing4 {
		thing4[i] = thing4[i] * thing4[i]
	}

	return *thing4
}
