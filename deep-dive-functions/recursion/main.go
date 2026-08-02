package main

import "fmt"

func main() {
	fact := factorial_loop(5)
	fmt.Println(fact)

	fact2 := factorial(5)
	fmt.Println(fact2)
}

// factorial of 5 => 5 * 4 * 3 * 2 * 1 => 120

// solution without recursion (with for loop)
func factorial_loop(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result = result * i
	}

	return result
}

// solution with recursion
func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)

}
