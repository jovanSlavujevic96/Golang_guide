package main

import "fmt"

func createNumbersArray() [10]int {
	numbers := [10]int{}
	for i := range numbers {
		numbers[i] = i
	}
	return numbers
}

func OddOrEven(number int) string {
	if number%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func main() {
	fmt.Print("Even or odd check\n\n")

	numbers := createNumbersArray()

	for _, num := range numbers {
		fmt.Println("Number ", num, " is ", OddOrEven(num), " number.")
	}
}
