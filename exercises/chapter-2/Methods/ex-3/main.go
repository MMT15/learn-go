package main

import "fmt"

type Celsius float64

func (c Celsius) ToFahrenheit() Celsius {
	F:=c* 9/5 +32
	return F
}

func (c *Celsius) IncreaseBy(delta float64) {
	*c+=Celsius(delta)
}

func main() {
	var tempC Celsius = 25.0
	fmt.Println(tempC.ToFahrenheit())
	fmt.Println(tempC)
	tempC.IncreaseBy(5.0)
	fmt.Println(tempC)
	
}
