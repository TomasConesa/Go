package main

import (
	"fmt"
)

func main() {

	myArrayVar := [3]int{5, 10, 11}
	fmt.Println("Array:", myArrayVar, "- size: ", len(myArrayVar))

	mySliceVar := []int{}
	mySliceVar = append(mySliceVar, 88, 99, 77)
	fmt.Println("Slice:", mySliceVar, "- size: ", len(mySliceVar))

	sliceVar2 := mySliceVar[1:2] // Puede ser [:3] y eso me toma desde el indice 0, lo mismo para el final [2:]
	fmt.Println("Slice:", sliceVar2, "- size: ", len(sliceVar2))

	// Inicializar el slice con make
	sliceVar3 := make([]int, 5)
	fmt.Println("Slice:", sliceVar3, "- size: ", len(sliceVar3))
}
