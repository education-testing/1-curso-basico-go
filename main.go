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
	print(isPalindrome(word))
}

func isPalindrome(word string) bool {
	word = strings.ToLower(word) //Normalizar la información
	var newWord string
	for i := len(word) - 1; i >= 0; i-- {
		newWord += string(word[i])
	}

	return word == newWord
}
