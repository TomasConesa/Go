package main

import (
	"encoding/json"
	"fmt"

	"github.com/TomasConesa/Go/structs/commerce"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Address  string `json:"address,omitempty"` // Si quiero que se omita un valor vacio en el json
	Age      uint8  `json:"age"`
}

func (u User) Display() {
	v, _ := json.Marshal(u)
	fmt.Println(string(v))
}

func (u *User) SetName(name string) {
	u.Name = name
}

func main() {
	user := User{
		Id:   1,
		Name: "Tomas",
	}
	user.Display()
	user.SetName("toteir00")
	user.Display()

	p1 := commerce.Product{
		Name:  "Heladera",
		Price: 1000,
		Type: commerce.Type{
			Code:        "A",
			Description: "Electrodomestico",
		},
		Tags:  []string{"heladera", "freezer", "refrigerador"},
		Count: 5,
	}

	p2 := commerce.Product{
		Name:  "Auriculares",
		Price: 500,
		Type: commerce.Type{
			Code:        "B",
			Description: "Tecnología",
		},
		Tags:  []string{"audio", "auricular", "sonido"},
		Count: 2,
	}

	p3 := commerce.Product{
		Name:  "Televisor",
		Price: 3000,
		Type: commerce.Type{
			Code:        "B",
			Description: "Tecnología",
		},
		Tags:  []string{"pantalla hd", "televisor", "sonido"},
		Count: 2,
	}

	cart := commerce.NewCart(10)
	cart.AddProducts(p1, p2, p3)

	fmt.Println("Productos cart")
	fmt.Println("Total products:", len(cart.Products))
	fmt.Printf("Total cart %.2f \n", cart.Total())
}
