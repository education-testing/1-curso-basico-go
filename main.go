package main

import (
	"fmt"
	"strings"
)

func main() {
	slice := []string{"Luis", "pedro", "Carlos"}

	for i, value := range slice {
		fmt.Println("Indice: ", i, "Value: ", value)
	}

	word := "amor a roma"
	fmt.Println(isPalindrome(strings.ToLower(word)))
}

func isPalindrome(word string) bool {
	var result string
	for i := len(word) - 1; i >= 0; i-- {
		result += string(word[i])
	}

	return result == word
}
