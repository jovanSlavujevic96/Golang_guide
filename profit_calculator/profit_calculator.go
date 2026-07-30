package main

import "fmt"

func main() {
	var revenue, expenses, tax_rate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax rate: ")
	fmt.Scan(&tax_rate)

	ebt := revenue - expenses
	profit := ebt * (1 - tax_rate/100)
	ratio := ebt / profit

	fmt.Println("EBT:", ebt)
	fmt.Println("profit:", profit)
	fmt.Println("ratio:", ratio)
}
