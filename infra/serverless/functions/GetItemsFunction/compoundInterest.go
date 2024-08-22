package main

func CalculateCompoundInterest(principal float64, annualRate float64, years int) []float64 {
	rateDecimal := annualRate / 100
	amountsPerYear := make([]float64, years)

	for year := 1; year <= years; year++ {
		principal = principal * (1 + rateDecimal)
		amountsPerYear[year-1] = principal
	}

	return amountsPerYear
}
