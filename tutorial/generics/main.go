package main

import "fmt"

type gasEngine struct {
	gallons float32
	mpg float32
}

type electricEngine struct {
	kwh float32
	mpkwh float32
}

type car [T gasEngine | electricEngine] struct {
	carMake string 
	carModel string
	engine T
}

func main () {
	// ---- First Example ----
	// var intSlice = []int{1,2,3}
	// // fmt.Println(sumSlice[int](intSlice))
	// fmt.Println(sumSlice(intSlice))

	// var floatSlice = []float32{1.5,2,3} 
	// // fmt.Println(sumSlice[float32](floatSlice))
	// fmt.Println(sumSlice(floatSlice))

	// --- Second Example ----
	var gasCar = car[gasEngine] {
		carMake: "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons: 12.4,
			mpg: 40,
		},
	}

	var electricCar = car[electricEngine] {
		carMake: "Tesla",
		carModel: "Model Y",
		engine: electricEngine{
			kwh: 57.5,
			mpkwh: 4.17,
		},
	}

	fmt.Println(gasCar)
	fmt.Println(electricCar)
}

func sumSlice[T int | float32 | float64] (slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}