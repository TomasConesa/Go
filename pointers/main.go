package main

import (
	"fmt"
)

func Increment(val *int) {
	*val++ // Incrementa el valor al que apunta el puntero
	fmt.Println("Inside Increment - Address:", val, "Value:", *val)
}

func updateSlice(s []int) {
	fmt.Printf("Address: %p \n", &s)
	fmt.Printf("Address0: %p, Address1: %p, Address2: %p \n", &s[0], &s[1], &s[2])
	s[0] = 10

}

type MyStruct struct {
	Id   int
	Name string
}

func (m MyStruct) Set(name string) {
	fmt.Printf("Address: %p \n", &m)
	m.Name = name
}

func (m *MyStruct) SetP(name string) {
	fmt.Printf("Address: %p \n", m)
	m.Name = name
}

func updateStruct(v *MyStruct) {
	fmt.Printf("Address in function: %p \n", v)
	v.Id = 999
	v.Name = "Tomas"
}

func main() {
	var i int
	var iP *int

	i = 34
	iP = &i

	fmt.Println(&i)
	fmt.Println(iP)
	fmt.Println(i)
	fmt.Println(*iP)

	*iP = 10
	fmt.Printf("Value i: %d, Value pointer: %d \n", i, *iP)

	myVar := 30
	fmt.Printf("Before increment address: %p, Value: %d \n", &myVar, myVar)
	Increment(&myVar)
	fmt.Printf("After increment address: %p, Value: %d \n", &myVar, myVar)

	var mySlice []int
	mySlice = append(mySlice, 5, 7, 8)
	fmt.Printf("Address: %p, Values: %v \n", &mySlice, mySlice)
	fmt.Printf("Address0: %p, Address1: %p, Address2: %p \n", &mySlice[0], &mySlice[1], &mySlice[2])
	updateSlice(mySlice)
	fmt.Println(mySlice)

	myStruct := &MyStruct{Id: 9, Name: "toteir00"} // Al declarar la variable de esta forma, ya indico que es un puntero de mi estructura
	fmt.Printf("Address: %p \n", myStruct)
	fmt.Printf("Id: %d, Name: %s \n", myStruct.Id, myStruct.Name)
	updateStruct(myStruct) // Como le paso un puntero, no hace falta indicarle & para saber la direccion en memoria
	fmt.Printf("Id: %d, Name: %s \n", myStruct.Id, myStruct.Name)
	fmt.Println() // Salto de linea

	// Pruebo los metodos sin utilizar puntero y utilizando puntero
	myStruct.Set("Test method")
	fmt.Printf("Id: %d, Name: %s \n", myStruct.Id, myStruct.Name)
	myStruct.SetP("Test method 2 pointer")
	fmt.Printf("Id: %d, Name: %s \n", myStruct.Id, myStruct.Name)
}
