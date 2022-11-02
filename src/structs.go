package main

import "fmt"

type car struct {
	brand string
	year int
}

func otherThink() {
	myCar := car{brand: "ford", year: 2020}


	fmt.Println(myCar)

	// Other form
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)

}