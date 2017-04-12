package main

import "fmt"

func main() {

	r := addProduct(Product{"kurtka", 500})

fmt.Println("Zawartość koszyka: ", r)

}


type Product struct {
	name string
	price int
}

func addProduct (product Product) [] Product{
	var (
		shoppingCart [] Product
	)
	shoppingCart = append( shoppingCart, product)
	return shoppingCart
}