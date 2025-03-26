package main 

import(
	"fmt"
)

func main(){

	yearsOld := 25

	// OR
	fmt.Println(yearsOld < 33 || yearsOld == 33) // true

	// AND
	fmt.Println(yearsOld < 33 && yearsOld == 33) // false

	// NOT
	fmt.Println(!(yearsOld == 25)) // false

	fmt.Println(yearsOld < 20 && yearsOld == 25 || yearsOld < 30) // false, true, true -> el resultado del && es false pero como se cumle el || va a ser true
	fmt.Println(yearsOld < 20 && (yearsOld == 25 || yearsOld < 30)) // false and true -> false



}