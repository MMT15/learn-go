package main

import "fmt"

var num1 int = 15
var num2 float64 = 4.5
var byteValue uint8 = 65
var runeValue rune = '😊'
var p,q int = 7,3
var counter int = 5

func main(){
	fmt.Print("+: ", float64(num1)+num2,
	"\n -: ", float64(num1)-num2,
	"\n *: ", float64(num1)*num2,
	"\n /: ", float64(num1)/num2,
	"\n")

	fmt.Printf("%c\n",byteValue)
	fmt.Printf("%c\n",runeValue)
	
	fmt.Print("&: ", p&q,
	"\n|: ", p|q,
	"\n<<(p): ", p<<1,
	"\n<<(q): ", q<<1,
	"\n")
	
	fmt.Println("counter initial: ", counter)
	counter++
	fmt.Println("counter++ = ", counter)
	counter--
	fmt.Println("counter-- = ", counter)
}
