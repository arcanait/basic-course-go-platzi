package mypackage

import "fmt"

type car struct {
	brand string
	year int
}

// CarPublic Car con acceso público
type CarPublic struct {
	Brand string
	Year int
}
// PrintMessage imprimir un mensaje
func PrintMessage(msg string) {
	fmt.Println(msg)
}

// PrintMessage imprimir un mensaje pero PRIVADA
func printMessage1(msg string) {
	fmt.Println(msg)
}