package main

import (
	"fmt"

	//"golang.org/x/exp/constraints"
	"cmp"
)

func main() {
	v1 := []float64{1.3, 5.5, 40.10, 55.55}
	v2 := []int32{7, 8, 9, 10, 11}

	/* 	fmt.Println(Sum(v1))
	   	fmt.Println(Sum(v2)) */

	fmt.Println(Sum01(v1))
	fmt.Println(Sum01(v2))

	anyType(5, 10)
	anyType2("5", 10)

	comparableType(1, 5)

	orderedValues(5, 10)

	csInt := CustomSlice[int]{1, 5, 8, 9}
	fmt.Println(csInt)
	csStg := CustomSlice[string]{"1", "5", "a", "b"}
	fmt.Println(csStg)

	fd := GenericStruct[MyFirstData]{StrValue: "Test", Data: MyFirstData{}}
	fd.Data.PrintOne()

	sd := GenericStruct[MySecondData]{StrValue: "Test", Data: MySecondData{}}
	sd.Data.PrintTwo()

}

func Sum[N int | int32 | int64 | float64 | float32](v []N) N {
	var total N
	for _, tV := range v {
		total += tV
	}
	return total
}

type Number interface {
	int | int32 | int64 | float64 | float32
}

func Sum01[N Number](v []N) N {
	var total N
	for _, tV := range v {
		total += tV
	}
	return total
}

func anyType[N any](v1, v2 N) {
	fmt.Println(v1, v2)
}

func anyType2[N1 any, N2 any](v1 N1, v2 N2) {
	fmt.Println(v1, v2)
}

func comparableType[N comparable](v1, v2 N) {
	fmt.Println(v1, v2)
	fmt.Println(v1 != v2)
}

func orderedValues[N cmp.Ordered](v1, v2 N) {
	fmt.Println(v1, v2)
	fmt.Println(v1 != v2)
	fmt.Println(v1 < v2)
	fmt.Println(v1 > v2)
}

type CustomSlice[V int | string] []V

func MinNumber[T Number01](x, y T) T {
	if x < y {
		return x
	}
	return y
}

type Point int

type Number01 interface {
	~int | ~int32 | ~int64 | ~float64 | ~float32
}

type (
	GenericStruct[D any] struct {
		StrValue string
		Data     D
	}

	MyFirstData struct{}

	MySecondData struct{}
)

func (d MyFirstData) PrintOne() {
	fmt.Println("Print first")
}

func (d MySecondData) PrintTwo() {
	fmt.Println("Print second")
}
