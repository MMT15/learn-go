package main

import "fmt"

var a,b int = 5,6

func greet(){
	fmt.Println("Hello, Go Functions!")
}

func add(a int, b int) int {
	c:=a+b
	return c
}

func substract(a,b int){
	fmt.Println(a-b)
}

func calc(a,b int) (int, int) {
	sum:=a+b
	prod:=a*b
	return sum,prod
}

func Calc(a,b int) (sum int, prod int) {
	sum=a+b
	prod=a*b
	return 
}

var op = func(a,b int) int {
	sum:=a+b
	return sum
}

func makeCounter() func() int {
	sum:=0
	return func () int {
		sum++
		return sum
	}
}

func sumAll(value ...int) int {
	sum:=0
	for _,v:=range value {
		sum+=v
	}
	return sum 
}

func init(){
	fmt.Println("Init function executed")
}

func init(){
	fmt.Println("Second Init function executed")
}

func cleanup(){
	fmt.Println("Cleanup finished!")
}

func main(){
	defer cleanup()

	greet()
	
	fmt.Println(add(a,b))
	substract(a,b)
	
	sum,prod:=calc(a,b)
	fmt.Println(sum, prod)
	
	Sum,Prod:=Calc(a,b)
	fmt.Println(Sum,Prod)
	
	fmt.Println(op(a,b))
	
	defer fmt.Println("It's finished?")
	
	func(){
		fmt.Println("This is an annonymous function")
	}()
	
	add:=makeCounter()
	fmt.Println(add())
	fmt.Println(add())
	fmt.Println(add())
	fmt.Println(add())

	fmt.Println(sumAll(1,2,3,4,5))
	fmt.Println(sumAll(10,20,30,40))
}
