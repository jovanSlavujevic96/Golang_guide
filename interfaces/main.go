package main

import "fmt"

// interfaces are not generic types
// interfaces are 'implicit'
// interfaces are a contract to help us manage types
// interfaces are tough. Step #1 is understanding how to read them
type bot interface {
	getGreeting() string
}

/**
* to whom it may concern...
* `type bot interface`
* Our program has a new type called `bot`
*
* `getGreeting() string`
* if you are a type in this program with a
* function called `getGreeting` and you return a
* string then you are now an honorary member
* of type `bot`
*
* Now that you're also an honorary member of
* type `bot`, you can now call this function
* called `printGreen`
* `func printGreeting(b bot)`
**/

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
