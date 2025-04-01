package main

import (
	"errors"
	"fmt"

	"github.com/TomasConesa/Go/errors/mypackage"
)

type MyCustomErr struct {
	Id string
}

func (m MyCustomErr) Error() string {
	return fmt.Sprintf("error with id: %s", m.Id)
}

func TestError() error {
	return MyCustomErr{"11"}
}

func TestError2(v int) error {
	if v == 1 {
		return MyCustomErr{"1"}
	} else {
		return errors.New("error")
	}
}

func main() {
	err := errors.New("new error")
	fmt.Println(err.Error()) // Esto es para formatear el error a tipo string

	err2 := fmt.Errorf("my format error, string: %s, number: %d", "toteir00", 10)
	fmt.Println(err2.Error())

	err3 := TestError()
	fmt.Println(err3)

	err4 := TestError2(1)
	fmt.Println(err4)

	err5 := TestError2(5)
	fmt.Println(err5)

	// Validar si el error esta implementado por la estructura

	err6 := TestError2(1)
	fmt.Println(errors.As(err6, &MyCustomErr{})) // true

	err7 := TestError2(5)
	fmt.Println(errors.As(err7, &MyCustomErr{})) // false

	// Join()
	err8 := errors.Join(err, err2, err3, err4, err6, err7)
	fmt.Println(err8)

	// Validar si contiene un error con Is()
	fmt.Println(errors.Is(err8, err))  // true
	fmt.Println(errors.Is(err8, err5)) // false

	// Anidando errores
	err9 := fmt.Errorf("error9: [%w]", err2)
	err10 := fmt.Errorf("error10: [%w]", err9)
	fmt.Println(err10)
	fmt.Println(errors.Unwrap(err10))                // Para que me muestr solo lo que está anidado
	fmt.Println(errors.Unwrap(errors.Unwrap(err10))) // Puedo hacerlo hasta que no tenga mas errores anidados y me devuelve <nil>

	// Defer, Panic y Recover

	defer func() {
		fmt.Println("end main")
		r := recover()
		if r != nil {
			fmt.Println("Recover in", r)
		}
	}()
	mypackage.Run()
	panic("my panic") // Fuerzo un error con panic, tiene que ir despues del Defer
}
