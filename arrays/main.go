package main 

import (
	"fmt"
	"unsafe"
)

func main() {
	var myIntVar int = 10
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", myIntVar, myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar) * 8)

	var arrayVar [5] int 
	fmt.Println(arrayVar, " - size: ", len(arrayVar))

	array2 := [2]string{"value1", "value2"}
	fmt.Println(array2, " - size:", len(array2)) 


	arrayVar[0] = 5
	arrayVar[1] = 10
	arrayVar[2] = 8
	arrayVar[3] = 7
	arrayVar[4] = 11
	fmt.Println(arrayVar, " - size: ", len(arrayVar))
	fmt.Printf("type: %T, value: %v, bytes: %d, bits: %d \n", arrayVar, arrayVar, unsafe.Sizeof(arrayVar), unsafe.Sizeof(arrayVar) * 8)

	// Recorrer un array
	for i := range arrayVar{
		fmt.Println("index: ", i, "- value:", arrayVar[i])
	}

	for i, v:= range arrayVar{
		fmt.Println("index: ", i, "- value:", v)
	}

}