package main 

import (
	"fmt"
)

func main() {

	map1 := make(map[int]string)
	map1[1] = "A"
	map1[2] = "B"
	map1[3] = "T"
	map1[10] = "J"

	fmt.Println(map1)
	fmt.Println(map1[3])

	
	map2 := make(map[int][]string)
	map2[1] = []string{"A", "B"}
	map2[5] = []string{"toteir00", "t0teiro"}
	fmt.Println(map2)
	fmt.Println(map2[5])

	map1[1] = "GG"
	fmt.Println(map1)

	// Eliminar una key
	delete(map1, 2)
	map1[9] = ""
	fmt.Println(map1)

	v, ok := map1[9]
	fmt.Println("El valor es: ", v, "- Existe:", ok)

 	v, ok = map1[8]
	fmt.Println("El valor es: ", v, "- Existe:", ok)

	map3 := map[int]string{
		1: "A",
		2: "B",
		3: "10",
	}
	fmt.Println(map3)

	// Recorrer un map

	for k := range map1 {
		fmt.Println("Key:", k, "- Value:", map1[k])
	}


	for key, value := range map2 {
		fmt.Println("Key:", key)

		// El _ es para que no me devuelva el index, que en este caso no lo quiero
		for _, v := range value {
			fmt.Println("Value:", v)
		}
	}
}