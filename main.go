// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"fmt"
)

/*
You can refer to official https://go.dev/ & https://pkg.go.dev/ site for reference of apis

Create a package products and create a type for Product as below
Product:
—----------
Name
Type : Grocery | Electronics | Kitchen
Price
Seller

1. Factory or initialize func to create Product
2. Add a method to apply discount to product. The method takes discount % as param and then applies discount to the price.
3. Add a method PrintJson that writes the json encoding of Product to console(os.Stdout)
use above methods in main for showcasing
In main package do following things 4 & 5 :
4. write Item interface with method to apply discount.
5. func ProcessItem(item Item). This takes an item iterface and applies 21.5% discount on item.
from main pass product to ProcessItem and print Product
*/

// Product docs
type Product struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Price  float32 `json:"price"`
	Seller string  `json:"seller"`
}

// Item
type Item interface {
	applyDiscount(percent float32)
}

// discount
func (p *Product) applyDiscount(percent float32) {
	p.Price = p.Price - (p.Price * percent / 100)
}

func ProcessItem(item Item) {
	item.applyDiscount(21.50) // apply discount 
}

func PrintJson(p Product) {
	jsonData, _ := json.Marshal(p)
	fmt.Printf("%s", string(jsonData))
}

func createProduct(name string, Type string, price float32, seller string) Product {
	var myProduct Product = Product{
		Name:   name,
		Type:   Type,
		Price:  price,
		Seller: seller,
	}
	return myProduct
}

// Main docs
func main() {
	var myProduct = createProduct("Mobile", "Electronics", 100.00, "Samsung") // create new product
	ProcessItem(&myProduct)                                                   // process item applies the discount
	PrintJson(myProduct)                                                      // print the product
}
