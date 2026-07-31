package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years, expectedReturnRate float64

	investmentAmount = getUserInput("Investment Amount: ")
	expectedReturnRate = getUserInput("Expected return rate: ")
	years = getUserInput("Years: ")

	futureValue, futureRealValue := calculateFinancial(investmentAmount, expectedReturnRate, years)

	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Value (adjusted for inflation):", futureRealValue)
	// fmt.Printf("Future Value: %.1f\nFuture Value (adjusted for inflation): %.1f\n)", futureValue, futureRealValue)

	formatedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formatedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.1f\n", futureRealValue)
	fmt.Print(formatedFV, formatedRFV)
}

func calculateFinancial(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	const inflationRate = 2.5
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv := fv / math.Pow(1+inflationRate/100, years)
	return fv, frv
}

func getUserInput(text string) (ret float64) {
	fmt.Print(text)
	fmt.Scan(&ret)
	return ret
}
