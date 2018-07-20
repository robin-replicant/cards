package main

import "fmt"

func main() {
	/*colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}*/

	// an other way of declaring
	//var colors map[string]string

	// other way
	//colors := make(map[string]string)
	//colors["white"] = "#ffffff"

	/*colors := make(map[int]string)
	colors[10] = "#ffffff"
	fmt.Println(colors)

	delete(colors, 10)

	fmt.Println(colors)*/

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
