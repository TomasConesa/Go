package main

import(
	"fmt"
)

func main(){

	// IF

	yearsOld := 25

	if yearsOld > 18 {
		fmt.Printf("%d is greather than 18 \n", yearsOld)
	}

	boolVar := true 

	if boolVar {
		fmt.Println("Es verdadero")
	} else {
		fmt.Println("Es falso")
	}

	if value := true; value {
		fmt.Println("Es verdadero")
	}

	number := 1

	if number == 1 {
		fmt.Println("Uno")
	} else if number == 5 {
		fmt.Println("Cinco")
	} else if number == 10 {
		fmt.Println("Diez")
	} else {
		fmt.Println("Ningun numero es valido")
	}

	// SWITCH

	switch number {
	case 1:
		fmt.Println("Uno")
	case 2: 
		fmt.Println("Dos")
	case 3: 
		fmt.Println("Diez")
	default: 
		fmt.Println("Ninguno es diez")
	}

	switch number := 2; number {
	case 1:
		fmt.Println("Uno")
	case 2: 
		fmt.Println("Dos")
	case 3: 
		fmt.Println("Diez")
	default: 
		fmt.Println("Ninguno es diez")
	}

	switch {
		case number < 0:
			fmt.Println("Negativo")
		case number > 0:
			fmt.Println("Positivo")
		case number == 0:
			fmt.Println("Cero")
	}
}