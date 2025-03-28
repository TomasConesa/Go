package main

import (
	"fmt"
)

func Increment(val *int) {
	*val++ // Incrementa el valor al que apunta el puntero
	fmt.Println("Inside Increment - Address:", val, "Value:", *val)
}

func main() {
	/* 	var i int
	   	var iP *int

	   	i = 34
	   	iP = &i

	   	fmt.Println(&i)
	   	fmt.Println(iP)
	   	fmt.Println(i)
	   	fmt.Println(*iP)

	   	*iP = 10
	   	fmt.Printf("Value i: %d, Value pointer: %d \n", i, *iP) */

	myVar := 30
	fmt.Printf("Before increment address: %p, Value: %d \n", &myVar, myVar)
	Increment(&myVar)
	fmt.Printf("After increment address: %p, Value: %d \n", &myVar, myVar)

}
