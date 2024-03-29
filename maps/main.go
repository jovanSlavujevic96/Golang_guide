package main

import "fmt"

func main() {
	// var colors map[string]string // creates empty map

	// make is a builtin function
	// colors := make(map[string]string) // creates empty map, too

	// create map with following types
	// [string] -> string as a key
	// string   -> string as a value
	colors := map[string]string{ // add some pairs
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// add new pair
	// colors["white"] = "#ffffff"

	// print map
	// fmt.Println(colors)

	// colorsInts := make(map[int]string)
	// colorsInts[0xffffff] = "white"

	// delete pair from map
	// delete(colors, "white")

	printColorMap(colors)
}

func printColorMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
