package main

import "fmt"

type Person struct {
	FirstName string
	LastName string
	Age int
	IsStudent bool
}

func main() {
	var person1 Person
	fmt.Println(person1)
	var person2 = Person{FirstName: "Jane", LastName: "Doe", Age: 30, IsStudent: true}
	fmt.Println(person2)
	var person3 = Person{FirstName: "Bob", Age: 25}
	fmt.Println(person3)
	fmt.Println(person2.FirstName, person2.Age)
}
