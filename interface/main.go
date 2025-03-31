package main

import (
	"fmt"
)

func Display(value interface{}) {
	fmt.Println(value)
}

func main() {
	Display("toteir00")
	Display(true)
	Display(10)

}
