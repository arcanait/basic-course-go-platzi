package main

import "fmt"

func keywords() {
	//Defer
	defer fmt.Println("hola") //Esto lo va a llamar de ultimas
	defer fmt.Println("hola2.0") //Esto lo va a llamar de penultimas
	defer fmt.Println("hola3.0") //Esto lo va a llamar de ante penultimas
	fmt.Println("mundo")

	//Continue
	for i := 0; i < 10; i++ {
		fmt.Println("i:%d", i)
		if i == 2 {
			fmt.Println("es 2")
			continue //Se utiliza para que en caso que el código se rompa dentro del ciclo este siga normal
		}
		//break
		if i == 8 {
			fmt.Println("Terminando ciclo...")
			break
		}
	}
}