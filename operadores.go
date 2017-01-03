package main

import (
	"fmt"
)

func main() {
	a := 49
	b := 7

	fmt.Println("\t Operadores \n", "****************************************")
	fmt.Println(a, "\t + \t", b, "\t = \t", a+b)
	fmt.Println(a, "\t - \t", b, "\t = \t", a-b)
	fmt.Println(a, "\t * \t", b, "\t = \t", a*b)
	fmt.Println(a, "\t / \t", b, "\t = \t", a/b)
	fmt.Println(a, "\t % \t", b, "\t = \t", a%b)

	c := 7
	fmt.Println("\t Comparadores \n", "****************************************")
	fmt.Println(a, "\t == \t", b, "\t ->\t ", a == b)
	fmt.Println(b, "\t == \t", c, "\t ->\t ", b == c)
	fmt.Println(a, "\t != \t", b, "\t ->\t ", a != b)
	fmt.Println(c, "\t != \t", b, "\t ->\t ", c != b)
	fmt.Println(a, "\t < \t", b, "\t ->\t ", a < b)
	fmt.Println(a, "\t <= \t", b, "\t ->\t ", a <= b)
	fmt.Println(c, "\t <= \t", b, "\t ->\t ", c <= b)
	fmt.Println(a, "\t > \t", b, "\t ->\t ", a > b)
	fmt.Println(a, "\t >= \t", b, "\t ->\t ", a >= b)

	fmt.Println("\t Operadores de Asignacion \n", "****************************************")
	fmt.Println(a,"\t -> \t", a+=b)
}
