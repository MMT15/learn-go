package calculator_test

import (
	"gomod/calculator"
	"testing"
)
func TestMultiply(t *testing.T) {
	result:=calculator.Multiply(3,4)
	expected:=12

	if result!=expected {
		t.Fail()
	}
}
