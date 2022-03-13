package main

import "fmt"

func main() {
	// Declaración de constantes.
	const pi float32 = 3.14
	const pi2 = 3.1415
	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	//Declaración de variables de forma 1:
	base := 12
	fmt.Println(base)

	//Declaración de variables de forma 2:
	var altura int = 32
	fmt.Println(altura)

	//Declaración de variables de forma 3:
	var anchura int
	fmt.Println(anchura)

	var a int
	var b float32
	var c string
	var d bool

	fmt.Println(a, b, c, d)

}
