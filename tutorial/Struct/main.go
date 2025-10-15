package main

import "fmt"

// To define a struct we use the type keyword
type gasEgine struct {
	mpg uint8
	gallons uint8
	ownerName owner
}

type owner struct {
	name string
}

func (e gasEgine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

// func canMakeIt(e gasEgine, miles uint8) {
func canMakeIt(e engine, miles uint8) {
	if miles < e.milesLeft() {
		fmt.Println("You can make it")
	} else {
		fmt.Println("Need to fuel up first :)")
	}
}

// Interface
type engine interface {
	milesLeft() uint8
}

func main () {
	var myEngine gasEgine = gasEgine{mpg: 25, gallons: 10, ownerName: owner{name: "Manas"}}
	// or 
	var myEngine2 gasEgine = gasEgine{25, 10, owner{name: "Manas"}}
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerName.name)
	fmt.Println(myEngine2.mpg, myEngine2.gallons, myEngine2.ownerName)

	myEngine.ownerName.name = "Sarah"
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerName.name)

	// Blocl scoped struct - not usable again
	var newStruct = struct {
		mpg uint8
		gallons uint8
	}{mpg: 25, gallons: 10}
	fmt.Println(newStruct.mpg, newStruct.gallons)

	// Now dealing with functions and methods
	fmt.Printf("Total miles left in the tank: %v\n", myEngine.milesLeft())

	canMakeIt(myEngine, 100)
}