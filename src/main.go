package main

import "fmt"

func print() {
	helloMessage := "Hello"
	worldMessage := "World"

	fmt.Println(helloMessage, worldMessage)

	platzi := "Platzi"
	cursos := 500

	fmt.Printf("%s este es un string:", platzi, "%d este es un número:", cursos, "\n") //Este se muestra raro por que entonces los valores tienen que ir de ulimas
	fmt.Printf("este es un string: %s, y este es un número: %d", platzi, cursos)
	fmt.Printf("este es un string: %v, y este es un número: %v \n" , cursos, platzi)

	//Sprintf
	message := fmt.Sprintf("%s este es un string", platzi)
	fmt.Println(message, "\n")

	// Data Types:
	fmt.Printf("What is the data type helloMessage?: %T \n", helloMessage)
	fmt.Printf("What is the data type cursos?: %T \n", cursos)
}