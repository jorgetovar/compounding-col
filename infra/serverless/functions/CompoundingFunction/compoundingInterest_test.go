package main

import (
	"reflect"
	"testing"
)

func TestCalculateCompoundInterest(t *testing.T) {
	principal := 100000000.0
	annualRate := 10.0
	years := 10

	result := CalculateCompoundInterest(principal, annualRate, years)
	lastElement := round(result[len(result)-1], 2)

	expected := round(259374246.01, 2)
	if !reflect.DeepEqual(lastElement, expected) {
		t.Errorf("Expected %v, but got %v", expected, lastElement)
	}
}
