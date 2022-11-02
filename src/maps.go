package main

import "fmt"

func maps() {
	m := make([]string, 2, 10) //Make aquí crea un array  []string, lo inicializa con length = 2 y con capacidad de 10
	m[0] = "london"
	fmt.Printf("m %v \n", m)

	otherMap := make(map[string]int)
	otherMap["jose"] = 15
	otherMap["pepito"] = 16

	fmt.Println(otherMap)

	// Recorrer map
	for i, v := range otherMap {
		fmt.Println(i, v)
	}

	fmt.Println(otherMap["jose"])
	joseName, exist := otherMap["jose"]
	joseName2, exist := otherMap["joses"]
	fmt.Println(joseName, exist) // Como no existe pone el mismo valor que tiene el 0 value, en este caso es 0 por que está con "int", con exist, nos va a retornar si el valor existe o no
	fmt.Println(joseName2, exist) // Como no existe pone el mismo valor que tiene el 0 value, en este caso es 0 por que está con "int", con exist, nos va a retornar si el valor existe o no
}