package main

import "fmt"

func main() {
	slice := []string{"Luis", "pedro", "Carlos"}

	for i, value := range slice {
		fmt.Println("Indice: ", i, "Value: ", value)
	}
}
