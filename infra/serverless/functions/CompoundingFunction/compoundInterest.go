package main

import "math"

func round(value float64, precision int) float64 {
	scale := math.Pow(10, float64(precision))
	return math.Round(value*scale) / scale
}

func CalculateCompoundInterest(principal float64, annualRate float64, years int) []float64 {
	rateDecimal := annualRate / 100
	gainsPerYear := make([]float64, years)

	for year := 1; year <= years; year++ {
		principal = principal * (1 + rateDecimal)
		gainsPerYear[year-1] = round(principal, 2)
	}

	return gainsPerYear
}
