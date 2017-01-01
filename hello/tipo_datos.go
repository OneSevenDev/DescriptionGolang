package main

import (
	"fmt"
)

func main() {
	//Numeros enteros
	var entero8 int8   //Todos los enteros de 8-bit (-128 to 127)
	var entero16 int16 //Todos los enteros de 16-bit (-32768 to 32767)
	var entero32 int32 //Todos los enteros de 32-bit (-2147483648 to 2147483647)
	var entero64 int64 //Todos los enteros de 64-bit (-9223372036854775808 to 9223372036854775807)

	//Numeros enteros positivos
	var uentero8 uint8   //Todos los enteros postivivos de 8-bit (0 to 255)
	var uentero16 uint16 //Todos los enteros postivivos de 16-bit (0 to 65535)
	var uentero32 uint32 //Todos los enteros postivivos de 32-bit (0 to 4294967295)
	var uentero64 uint64 //Todos los enteros postivivos de 64-bit (0 to 18446744073709551615)

	//Alias
	var enteroByte byte //alias para uint8
	var enteroRune rune //alias para int32

	//Tipos dependientes de la plataforma
	var enteroUint uint       //32 o 64 bit
	var enteroInt int         //32 o 64 bit
	var enteroUintptr uintptr //Entero sin signo lo suficientemente grande para contener el valor de un puntero

	fmt.Println("Enteros positivos y negativos \n ===========================")
	fmt.Println(entero8, entero16, entero32, entero64)
	fmt.Println("Enteros positivos \n ===========================")
	fmt.Println(uentero8, uentero16, uentero32, uentero64)
	fmt.Println("Alias de variables \n ===========================")
	fmt.Println(enteroByte, enteroRune)
	fmt.Println("Depende de la plataforma donde corre el programa \n ==================================================")
	fmt.Println(enteroUint, enteroInt, enteroUintptr)
}
