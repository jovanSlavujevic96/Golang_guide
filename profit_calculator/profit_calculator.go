package main

import "fmt"

func main() {
	var revenue, expenses, tax_rate float64

	getUserInfo("Revenue: ", &revenue)
	getUserInfo("Expenses: ", &expenses)
	getUserInfo("Tax rate: ", &tax_rate)

	ebt, profit, ratio := calculateFinancials(revenue, expenses, tax_rate)

	fmt.Println("EBT:", ebt)
	fmt.Println("profit:", profit)
	fmt.Println("ratio:", ratio)
}

func getUserInfo(text string, ret *float64) {
	fmt.Print(text)
	fmt.Scan(ret)
}

func calculateFinancials(revenue, expenses, tax_rate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - tax_rate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
