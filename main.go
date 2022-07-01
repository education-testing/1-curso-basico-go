package main

import (
	"fmt"
	"sync"
)

func say(name string, group *sync.WaitGroup) {
	fmt.Println("Hello", name)
	group.Done()
}

func sayAnother(name string, channel chan<- string) {
	channel <- name // Meter información en el canal
}

// Probando la concurrencia en go
func main() {
	name := "Luis"
	var group sync.WaitGroup
	group.Add(1) // Se agrega 1 ya que soló voy a tener una goroutine
	fmt.Println("Hello", name)
	go say("Eduardo", &group) // No se alcanza a ejecutar antes que el hilo del main

	group.Wait()

	func(parameter string) {
		fmt.Println("GoodBye", parameter)
	}(name)
	//time.Sleep(time.Second * 1) // Esto lo usamos para tardar el hilo de main, y poder ver el hilo de say()

	/*
		Creación de los canales
	*/

	channel := make(chan string, 1)
	go sayAnother(name, channel)
	fmt.Println("mi canal", <-channel)

}
