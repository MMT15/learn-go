package main

import "fmt"

type Address struct {
	Street string
	City string
}

type Customer struct {
	Address
	CustomerID string
}

type Supplier struct {
	BillingAddress Address
	SupplierId string
}

func main() {
	customer1 := Customer{}
	customer1.CustomerID="John"
	customer1.Street="123 Main St"
	customer1.City="New York"
	
	supplier1 := Supplier{}		
	supplier1.SupplierId = "Petter Lmt"
	supplier1.BillingAddress.Street = "111 Main st"
	supplier1.BillingAddress.City = "Ohio"

	fmt.Println(customer1.Street, customer1.City)
	fmt.Println(supplier1.BillingAddress.Street,supplier1.BillingAddress.City)
	
}
