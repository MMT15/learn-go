package calculator_test

import (
	"gomod/calculator"
	"testing"
)

type multiplyTestCase struct {
	a int
	b int
	expected int
}

func TestMultiply(t *testing.T) {
	var tester = []multiplyTestCase{
		{3, 4, 12},
		{0, 5, 0},
		{5, 0, 0},
		{-5, -5, 25},
		{-5, 5, -25},
	}
	for _,v:= range tester {
		result:=calculator.Multiply(v.a,v.b)
		if result != v.expected {
			t.Errorf("For %d * %d, expected %d, got %d\n", v.a, v.b, v.expected, result)
		}
	}
}
