package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum_slice := sumup_slice(numbers)

	sum := sumup(1, 10, 15)

	fmt.Println(sum_slice)
	fmt.Println(sum)
}

func sumup_slice(numbers []int) int {
	sum := 0

	for _, val := range numbers {
		sum += val // sum = sum + val
	}

	return sum
}

// variadic func
func sumup(startingValue int, numbers ...int) int {
	sum := startingValue

	for _, val := range numbers {
		sum += val // sum = sum + val
	}

	return sum
}
