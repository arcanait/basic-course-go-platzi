package main

import "fmt"

func arrays() {
	// Array - Inmutable
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array)

	// Slice  - Mutable
	slice := []int{0,1,2,3,4,5,6}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos para slices
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append  - Agrega un elemento
	slice = append(slice, 7)  
	fmt.Println("slice: %v",slice)

	// Append nueva list
	newSlice := []int{8,9,10}
	slice = append(slice, newSlice...)
	fmt.Println("slice: %v",slice)

	// Recorriendo array
	otherSlice := [3]string{"hola", "que", "hace"}
	fmt.Println("otherSlice %s", otherSlice)
	for i, valor := range otherSlice {
		fmt.Printf("i: %d, valor: %v \n", i, valor) // Imprime el indice y el valor de cada uno de los elementos del array
	}

	for _, valor := range otherSlice {
		fmt.Printf("valor: %v \n", valor) // Imprime solo el valor de cada uno de los elementos
	}

	for valor := range otherSlice {
		fmt.Printf("valor: %v \n", valor) // Imprime el índice únicamente
	}
}