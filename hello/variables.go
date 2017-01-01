package main

import (
	"fmt"
)

var barra = "código de barras"

// barra := "código de barras" -> no permite, solo el := es valido dentro de funciones

func main() {

	//Varibles
	_number := 17             //enteros
	_name := "Manuel Guarniz" //cadena

	var (
		bol      = false
		integer  = 21
		textShow = "Texto de muestra"
	)

	fmt.Println(_number, _name, "\n ========================= \n", bol, integer, textShow)

	//Ejemplos de variables
	var number int
	number = 25
	fmt.Println("========================= \n", number)
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

	//scope
	//var barra = "código de barras" -> se cambio a variable global
	fmt.Println("Al reverso del producto esta el " + barra)
}

func printText() {
	fmt.Println("Funciona adicional: Al reverso del producto esta el " + barra)
}
