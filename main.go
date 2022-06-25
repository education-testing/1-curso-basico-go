package main

import (
	p "1-curso-basico-go/classes"
	fi "1-curso-basico-go/interfaces"
	"fmt"
)

func calculate(figures fi.Figures) {
	fmt.Println(figures.GetArea())
}

func main() {
	a := 50
	b := &a // Creo la referencia a la variable a

	*b = 500

	fmt.Println("b, es una referencia a a", b)         // Abro la puerta
	fmt.Println("El valor de a desde la referncia", a) //saco el valor

	myPc := p.Pc{Memory: 6, Cpu: 10, Brand: "lenovo"}
	fmt.Println(myPc)
	myPc.SetCpu(15)
	fmt.Println(myPc)

	mySquare := p.Square{Base: 45.3}
	myRectangle := p.Rectangle{Base: 45.3, Height: 58.2}

	fmt.Println(mySquare)
	fmt.Println(myRectangle)

	calculate(mySquare)
	calculate(myRectangle)

	//Declaración de una lusta de interfaces:
	interfaceLis := []interface{}{"Hello world", 4.5, 5, true}
	fmt.Println(interfaceLis)

	for i, v := range interfaceLis {
		fmt.Println(i, v)
	}

}
