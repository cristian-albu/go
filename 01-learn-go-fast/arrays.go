package main

import (
	"fmt"
	"time"
)

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

func Arrays() {
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

	slicesPerformance()
}
