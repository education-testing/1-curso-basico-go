package main

import (
	"fmt"
	"math"
)

func main() {
	//suma
	var x int = 5
	var y int = 6
	result := x + y
	fmt.Println("sum: ", result)

	//Resta
	result = y - x
	fmt.Println("subtract: ", result)

	//Multiplicación
	result = y * x
	fmt.Println("Multiply: ", result)

	//División
	result = y / x
	fmt.Println("Division: ", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo: ", result)

	//Incrementar
	x++
	fmt.Println("Incrementar x: ", x)

	//Decrementar
	x--
	fmt.Println("Decrementar: ", x)

	var altura int = 5
	var longitudA int = 6
	var longitudB int = 7
	var radio int = 9

	areaTrapecio := altura * ((longitudA + longitudB) / 2)
	areaRectangulo := altura * longitudA
	areaCirculo := math.Pi * math.Exp2(float64(radio))

	fmt.Println("Area Trapecio: ", areaTrapecio)
	fmt.Println("Area Rectangulo : ", areaRectangulo)
	fmt.Println("Area Circulo : ", areaCirculo)

}
