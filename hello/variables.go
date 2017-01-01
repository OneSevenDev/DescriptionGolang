package main

import (
	"fmt"
)

func main() {

	var number int
	number = 25
	fmt.Println(number)
	number = 40
	fmt.Println(number)

	name := "Manuel" // or var name string
	fmt.Println(name)

	//Asignacion de variables en orden
	name, number = "Esteban", 21
	fmt.Println(name, number)

	//cambio de valores entre variables
	lastName := "Guarniz"
	fmt.Println(name, lastName)
	name, lastName = lastName, name
	fmt.Println(name, lastName)
}
