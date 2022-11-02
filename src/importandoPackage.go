package main

import (
	"fmt"
	pk "src/mypackage/src/mypackage"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	fmt.Println(myCar)

	pk.PrintMessage("holaaaa")
}