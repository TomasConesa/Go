package commerce

type Type struct {
	Id          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Product struct {
	Id    uint     `json:"id"`
	Name  string   `json:"name"`
	Type  Type     `json:"type"`
	Count uint16   `json:"count"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

type Cart struct {
	Id       uint      `json:"id"`
	UserId   uint      `json:"user_id"`
	Products []Product `json:"products"`
}

func (p Product) TotalPrice() float64 {
	return float64(p.Count) * p.Price
}

func (c *Cart) AddProducts(products ...Product) {
	c.Products = append(c.Products, products...)
}

func (c Cart) Total() float64 {
	var total float64
	for _, p := range c.Products { // Con el _ omito el indice porque en este caso no lo necesito, pero es necesario para recorrer el slice
		total += p.TotalPrice()
	}
	return total
}

func NewCart(userId uint) Cart {
	return Cart{UserId: userId}
}
