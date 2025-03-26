package function 

import(
	"fmt"
	"errors"
)

type Operation int
const (
	SUM Operation = iota 
	SUB
	DIV
	MUL 
)

// Funciones declaradas con mayúscula al principio, son públicas. En minúscula, son privadas y no se pueden acceder desde otro package

func Display(value int64){
	fmt.Printf("Type: %T, Value: %d, Address: %v \n", value, value, &value)
}

func Add(x int, y int) int{
	return x + y
}

func RepeatString(increment int, value string){
	for i := 0; i < increment; i++{
		fmt.Println(value)
	}
}

func Calc(op Operation, x, y float64)(float64, error){
	switch op {
	case SUM:
		return x + y, nil
	case SUB:
		return x - y, nil 
	case DIV:
		if y == 0 {
			return 0, errors.New("Y musn't be zero")
		}
		return x / y, nil 
	case MUL:
		return x * y, nil
	}

	return 0, errors.New("Has been an error")
}

func Split(v int) (x, y int){
	x = v * 5/10
	y = v - x
	return 
}

// El ... solo se puede usar una vez dentro de los parámetros y tiene que ir al final, en caso de tener mas de 1 parámetro
func MSum(values ...float64) float64{
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum 
}

func MOperations(op Operation, values ...float64) (float64, error){
	if len(values) == 0 {
		return 0, errors.New("There aren't values")
	}
	sum := values[0]

	for _, v := range values[1:] {
		switch op {
		case SUM :
			sum += v 
		case SUB:
			sum -= v 
		case DIV: 
			if v == 0 {
				return 0, errors.New("V musn't be zero")
			}
			sum /= v 
		case MUL:
			sum *= v
		}
	}
	return sum, nil 
}