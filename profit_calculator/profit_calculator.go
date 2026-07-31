package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var revenue, expenses, tax_rate float64
	var e error

	revenue, e = getUserInfo("Revenue: ")
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}
	expenses, e = getUserInfo("Expenses: ")
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}
	tax_rate, e = getUserInfo("Tax rate: ")
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, tax_rate)

	fmt.Println("EBT:", ebt)
	fmt.Println("profit:", profit)
	fmt.Println("ratio:", ratio)
	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	resutls := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(resutls), 0644)
}

func isInputValid(input float64) bool {
	// no negative numbers
	// no zero value
	// therefore input must be greater than zero
	return (input > 0.0)
}

func getUserInfo(text string) (float64, error) {
	fmt.Print(text)

	var ret float64
	fmt.Scan(&ret)

	if isInputValid(ret) {
		return ret, nil
	}
	return ret, errors.New("Invalid value for " + text + strconv.FormatFloat(ret, 'f', -1, 64))
}

func calculateFinancials(revenue, expenses, tax_rate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - tax_rate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
