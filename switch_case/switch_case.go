package switch_case

import (
	"fmt"
)

// No break is needed go automatically adjust the break statement
// You can match multiple values in one case (case "a","b":).
// default: works like else.
func Switch() {
	var color string
	fmt.Println("Please enter color")
	fmt.Scanln(&color)
	switch color {
	case "red":
		fmt.Println("You picked red")
	case "green":
		fmt.Println("You picked green")
	case "blue", "purple":
		fmt.Println("Something of blue and purple")
	default:
		fmt.Println("You choice didn't matched")

	}
}
