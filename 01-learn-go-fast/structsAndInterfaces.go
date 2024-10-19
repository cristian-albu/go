package main

import "fmt"

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

func StructsAndInterfaces() {

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
