package main

import "fmt"

func modify(name string) string {
	return name + "hello"
}

func receiveMessage(channel chan<- string, message string) {
	channel <- modify(message)
}

func main() {
	channel := make(chan string, 2)
	channel <- "Mensaje1"
	channel <- "Mensaje2"
	fmt.Println(len(channel))
	fmt.Println(cap(channel))

	close(channel)

	for message := range channel {
		fmt.Println(message)
	}

	email1 := make(chan string)
	email2 := make(chan string)

	go receiveMessage(email1, "Email1")
	go receiveMessage(email2, "Email2")
}
