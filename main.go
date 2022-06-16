package main

import (
	c "1-curso-basico-go/my-package" // usando alias
	"fmt"
)

func main() {
	car := c.Car{Brand: "Chevrolet", Year: 2022} //Clase publica
	fmt.Println(car)

	fmt.Println(c.HelloWord()) // Metodo publico
}
