package main

import (
	"fmt"
	"unsafe"
	"strconv"
)

func main() {
	var myIntVar int
	myIntVar = -15
	fmt.Println("My variable is:", myIntVar)

	// uint solo acepta valores positivos
	var myUnitVar uint
	myUnitVar = 15
	fmt.Println("My variable is:", myUnitVar)

	var myStringVar string 
	myStringVar = "Toteiro"
	fmt.Println("My variable is:", myStringVar)

	var myBoolVar bool
	myBoolVar = true
	fmt.Println("My variable is:", myBoolVar)

	// Dirección en memoria de variable con &
	fmt.Println("My variable address is:", &myStringVar)

	// Declarar y asignarle un valor a la variable con :=
	intVar2 := 10
	fmt.Println("My variable is:", intVar2)

	myStringVar2 := "EDLP"
	fmt.Println("My variable is:", myStringVar2)

	myBoolVar2 := true 
	fmt.Println("My variable is:", myBoolVar2)

	// CONSTANTES -> valores que no se pueden modificar, van a ser solo de lectura

	const myIntConst int = 99
	fmt.Println("My const is:", myIntConst)

	const myStringConst = "La scaloneta"
	fmt.Println("My const is:", myStringConst)
 
	// BLOQUES

	myOtherScope := 50

	{
		myFirstScope := 88
		{
			fmt.Println("My variable:", myFirstScope)
			fmt.Println("My variable:", myOtherScope)
		}
		fmt.Println("My variable:", myFirstScope)
		fmt.Println("My variable:", myOtherScope)
	}

	// fmt.Println("My variable:", myFirstScope) -> No puedo acceder


	// VARIABLE SIZE

	// Valores para int y uint

	// int8   - 8 bits   - -128 to 127
	// uint8  - 8 bits   - 0 to 255

	// int16  - 16 bits  - -32,768 to 32,767
	// uint16 - 16 bits  - 0 to 65,535

	// int32  - 32 bits  - -2,147,483,648 to 2,147,483,647
	// uint32 - 32 bits  - 0 to 4,294,967,295

	// int64  - 64 bits  - -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
	// uint64 - 64 bits  - 0 to 18,446,744,073,709,551,615

	// int    - Depende de la arquitectura del sistema (32 o 64 bits)
	// uint   - Depende de la arquitectura del sistema (32 o 64 bits)

	var my8BitsUintVar uint8 = 20
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my8BitsUintVar, my8BitsUintVar, unsafe.Sizeof(my8BitsUintVar), unsafe.Sizeof(my8BitsUintVar) * 8)
	
	var my16BitsUintVar uint16 = 30
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my16BitsUintVar, my16BitsUintVar, unsafe.Sizeof(my16BitsUintVar), unsafe.Sizeof(my16BitsUintVar) * 8)

	var my32BitsUintVar uint32 = 90
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my32BitsUintVar, my32BitsUintVar, unsafe.Sizeof(my32BitsUintVar), unsafe.Sizeof(my32BitsUintVar) * 8)	

	var my64BitsUintVar uint64 = 90
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my64BitsUintVar, my64BitsUintVar, unsafe.Sizeof(my64BitsUintVar), unsafe.Sizeof(my64BitsUintVar) * 8)	

	var uintVar uint = 90
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", uintVar, uintVar, unsafe.Sizeof(uintVar), unsafe.Sizeof(uintVar) * 8)	 */

	// FLOAT

 	var float32Var float32 = 3.14
	fmt.Printf("type: %T, value: %f, bytes: %d, bits: %d \n", float32Var, float32Var, unsafe.Sizeof(float32Var), unsafe.Sizeof(float32Var) * 8)

	var float64Var float64 = 590.12345
	fmt.Printf("type: %T, value: %f, bytes: %d, bits: %d \n", float64Var, float64Var, unsafe.Sizeof(float64Var), unsafe.Sizeof(float64Var) * 8)

	myOtherFloat := 7000.1253
	fmt.Printf("type: %T, value: %f, bytes: %d, bits: %d \n", myOtherFloat, myOtherFloat, unsafe.Sizeof(myOtherFloat), unsafe.Sizeof(myOtherFloat) * 8) */

	// STRINGS

	var stringVar string = "totoconesa10"
	fmt.Printf("Mi valor %s \n", stringVar)

	stringVar2 := `Variable de tipo texto en go 
con multiples
lineas
:)`
	fmt.Printf("Mi valor es %s \n", stringVar2) 

	// CONVERT TO STRING

	{
		floatVar := 33.11
		fmt.Printf("type: %T, value: %f \n", floatVar, floatVar)

		floatStrVar := fmt.Sprintf("%.2f", floatVar)
		fmt.Printf("type: %T, value: %s \n", floatStrVar, floatStrVar)

		intVar3 := 10
		intStrVar3 := strconv.Itoa(intVar3)
		fmt.Printf("type: %T, value: %s \n", intStrVar3, intStrVar3)
	} */

	// CONVERT TO NUMBER

	strIntVar, err := strconv.Atoi("15")
	fmt.Printf("type: %T, value: %d, err: %v \n", strIntVar, strIntVar, err)

	strIntVar2, err := strconv.ParseInt("10", 10, 64)
	fmt.Printf("type: %T, value: %d, err: %v \n",strIntVar2, strIntVar2, err)

	strIntVar3, err := strconv.ParseFloat("10.55", 64)
	fmt.Printf("type: %T, value: %.2	f, err: %v \n",strIntVar3, strIntVar3, err)


}
