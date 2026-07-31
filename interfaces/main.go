package main

import "fmt"

// interfaces are not generic types
// interfaces are 'implicit'
// interfaces are a contract to help us manage types
// interfaces are tough. Step #1 is understanding how to read them
type bot interface {
	getGreeting() string
}

// embedded interface
// in order to be an interface struct must implement both `bot` & `botPlus` methods
type botPlus interface {
	bot
	getParting() string
}

// interface{} type means func receives any kind of value/type for an argument
// option 1 with switch-case
// func printSomething(value interface{}) {
// 	switch value.(type) { // different handling between types
// 	case int:
// 		fmt.Println("Integer:", value)
// 	case float64:
// 		fmt.Println("Float:", value)
// 	case string:
// 		fmt.Println(value)
// 	default:
// 		// don't do anything for any other type of value
// 	}
// }

// option 2 with extracting type
// func printSomething(value interface{}) {
// 	intVal, ok := value.(int)
// 	if ok {
// 		fmt.Println("Integer:", intVal)
// 		return
// 	}

// 	floatVal, ok := value.(float64)
// 	if ok {
// 		fmt.Println("Float:", floatVal)
// 		return
// 	}

// 	stringVal, ok := value.(string)
// 	if ok {
// 		fmt.Println("String:", stringVal)
// 		return
// 	}
// }

// option 3 with generics
func printSomething[T int | float64 | string /*any*/](value T) {
	fmt.Println(value)
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

	printGreetingAndParting(eb)
	printGreetingAndParting(sb)

}

func printGreeting(b bot) {
	printSomething(b.getGreeting())
}

func printParting(b botPlus) {
	printSomething(b.getParting())
}

func printGreetingAndParting(b botPlus) {
	printGreeting(b)
	printParting(b)
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

func (englishBot) getParting() string {
	return "Bye"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

func (spanishBot) getParting() string {
	return "Adios"
}
