package main

import (
	"fmt"
	"strings"
)

func main() {
	sliceString := []string{"Hola", "que", "hace"}

	for i, _ := range sliceString {
		fmt.Println(i)
	}

	word := "amor a roma"
	isPalindrome(word)

	//// Mapas
	m := make(map[string]int)
	m["hector"] = 12
	m["Carlos"] = 21

	for key, value := range m {
		fmt.Println(key, value)
	}

	//Encontrando un valor
	value, ok := m["hectorino"]
	fmt.Println(ok, value)

}

func isPalindrome(word string) bool {
	word = strings.ToLower(word) //Normalizar la información
	var newWord string
	for i := len(word) - 1; i >= 0; i-- {
		newWord += string(word[i])
	}

	return word == newWord
}
