package main

import "fmt"

func main() {
	// create map with following types
	// [string] -> string as a key
	// string   -> string as a value
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	fmt.Println(colors)
}
