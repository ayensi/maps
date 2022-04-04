package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"white": "#FFFFFF",
		"black": "#000000",
	}
	/*var colors map[string]string*/
	//colors := make(map[string]string)
	printMap(colors)

}
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%s's hex code is: %v \n", color, hex)
	}
}
