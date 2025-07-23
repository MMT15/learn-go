package main

import "fmt"

type Product struct {
	Name string
	Price float64
}

func updateProductValue(prod Product) {
	prod.Price+=200
}

func updateProductPointer(prod *Product) {
	prod.Price+=200
}

func main() {
	var item = Product{Name: "Laptop",Price: 1200.0}
	fmt.Println(item)
	updateProductValue(item)
	fmt.Println(item)
	updateProductPointer(&item)
	fmt.Println(item)
}