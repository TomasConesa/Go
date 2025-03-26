package main

import(
	"fmt"
	"github.com/TomasConesa/Go/functions/function"
)

func main () {
 	var intVar int64
	var intVar2 int64 = 10
	var intVar3 int64 = 7

	function.Display(intVar)
	function.Display(intVar2)
	function.Display(intVar3)

	fmt.Println() // Salto de linea
	resultadoSuma := function.Add(5,5)
	fmt.Println(resultadoSuma)

	function.RepeatString(5,"toteir00")

	valueCalc, error := function.Calc(function.DIV, 33.3, 0)
	fmt.Println("Value:", valueCalc, "- Error:", error)

	xVal, yVal := function.Split(50)
	fmt.Println(xVal, yVal)

	val2 := function.MSum(1,10,15,4898)
	fmt.Println(val2)
 

 	mOperVal, err := function.MOperations(function.SUM, 10,10,100,77,99)
	fmt.Println("Value:",mOperVal, "- Error:", err)

	
	mOperVal, err = function.MOperations(function.SUB, 10,10,100,77,99)
	fmt.Println("Value:",mOperVal, "- Error:", err)


	mOperVal, err = function.MOperations(function.DIV, 10,0,100,77,99)
	fmt.Println("Value:",mOperVal, "- Error:", err)


	mOperVal, err = function.MOperations(function.MUL)
	fmt.Println("Value:",mOperVal, "- Error:", err) 

	factOpFunc := function.FactoryOperation(function.SUM)
	fmt.Println(factOpFunc(10,8))

	factOpFunc = function.FactoryOperation(function.SUB)
	fmt.Println(factOpFunc(10,8))

	factOpFunc = function.FactoryOperation(function.MUL)
	fmt.Println(factOpFunc(10,8))

	factOpFunc = function.FactoryOperation(function.DIV)
	fmt.Println(factOpFunc(10,8))


}
