package main

import "fmt"

func main() {
	//Los arrays son immutables.
	var array [7]int8
	array[0] = 32
	array[6] = 1

	tamano := len(array)
	capacidad := cap(array)

	fmt.Println("Tamaño del array: ", tamano)
	fmt.Println("Capacidad del array: ", capacidad)
	fmt.Println(array)

	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice)
	fmt.Println("len(): ", len(slice))
	fmt.Println("cap(): ", cap(slice))

	/*Jugando con los slices y los arrays*/

	//Imprimiendo la primer posición
	fmt.Println(array[0])
	fmt.Println(slice[0])

	//Imprimiendo hasta la posición 3
	fmt.Println(array[:3])
	fmt.Println(slice[:3])

	//Imprimiendo los elementos del indice 2 la indice 4
	fmt.Println(array[2:4])
	fmt.Println(slice[2:4])

	//Imprimiendo del indice 4 en adelante.
	fmt.Println(array[4:])
	fmt.Println(slice[4:])

	//Agregando un nuevo valor :D
	slice = append(slice, 7)
	fmt.Println(slice)

	//Creamos un nuevo slice para agregarlo en el original.
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10)
}
