package main

import "fmt"

func main() {
	// new map with string keys and string values
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"blue":  "#0000ff",
		"white": "#FFFFFF",
	}

	// var colors map[string]string

	/*
		colors := make(map[string]string)

		colors["white"] = "#FFFFFF"

		fmt.Println(colors)
	*/

	/*
		cars := make(map[int]string)

		cars[10] = "Audi"
		cars[9] = "Ford"

		fmt.Println(cars)
		delete(cars, 9)
		fmt.Println(cars)
	*/
	printMap(colors)
}

func printMap(m map[string]string) {
	for colorKey, colorValue := range m {
		fmt.Println("Hex code for", colorKey, "is", colorValue)
	}
}
