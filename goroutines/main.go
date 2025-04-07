package main

import (
	"fmt"
	"time"
)

func main() {
	go MyProcess("A")
	go MyProcess("B")
	time.Sleep(3 * time.Second) // Cuando se cumplen los 3 segundos el main finaliza y los procesos que queden por correr terminan y no se ejecutan mas

	myFirstChannel := make(chan string)        // Genero el canal
	go ProcessWithChannel("C", myFirstChannel) // Debo ejecutarlo como una goroutine
	result := <-myFirstChannel                 // Recibo el valor del canal en una variable
	fmt.Println(result)
	close(myFirstChannel) // Debo cerrar el canal

	channelA := make(chan string)
	channelB := make(chan string)

	go ProcessWithChannel("D", channelA)
	go MyOtherProcessWithChannel("E", channelB)

	resultA := <-channelA
	fmt.Println("A: ", resultA)
	resultB := <-channelB
	fmt.Println("B: ", resultB)
}

// Creo una funcion que recibe un string y el canal
func ProcessWithChannel(p string, c chan string) {
	x := 0
	for x < 5 {
		x++
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("Process: %s - Num: %d \n", p, x)
	}
	c <- "ok" // Asigno un valor al canal
}

func MyOtherProcessWithChannel(p string, c chan string) {
	x := 0
	for x < 20 {
		x++
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("Process: %s - Num: %d \n", p, x)
	}
	c <- "ok" // Asigno un valor al canal
}

func MyProcess(p string) {
	i := 0
	for i < 15 {
		i++
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("Process: %s - Num: %d \n", p, i)
	}
}
