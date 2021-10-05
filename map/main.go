package main

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4b745",
		"white": "#fffff",
	}

	// colors["white"] = "#fffff"

	// fmt.Println(colors)
	// fmt.Println(colors["white"])

	// delete(colors, "white")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		println(key, value)
	}
}