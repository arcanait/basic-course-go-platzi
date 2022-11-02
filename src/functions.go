package main

import (
	"fmt"
)

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b ,c)
}

func returnValue(a int) int {
	return a * 2
}

func returnDoubleValue(a int) (b, c int) {
	return a, a*2
}

func functions() {
	messageList := []string{"hola", "mundo", "soy", "yo"}
	numberList := []int{1,2,3,4,5}

	for i := 0; i < len(messageList); i++ {
		normalFunction(messageList[i])
	}
	for i := 0; i < len(numberList); i++ {
		fmt.Printf("Este es el número: %d", returnValue(numberList[i]))
	}

	for i := 0; i < len(numberList); i++ {
		value1, value2 := returnDoubleValue(numberList[i])
		message := fmt.Sprintf("Este es el número: %d y este es el número al multiplicarlo por 2 da: %d", value1, value2)
		fmt.Println(message)
	}

}