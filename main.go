package main

import (
	"firstProject/json_handling"
	"fmt"
)

func main() {
	fmt.Println("Hello World")                    //"Hello world\n"
	fmt.Print("Hello")                            //"Hello"
	fmt.Printf("Name: %s, Age: %d", "Prince", 20) // Using %s for a string and %d for an integer
	val := fmt.Sprint("The answer is", 42)        //To make value in string
	fmt.Println(val)

	/******Scan Property which take input from user*****/
	// var name string
	// var prompt string
	// fmt.Println("Please Enter your name and prompt")
	// fmt.Scan(&name, &prompt) //it take input with respect of one space
	// fmt.Println(prompt, name)
	// fmt.Scanln(&name, &prompt)
	// println(prompt, name) // it will make space by default

	// data_types.ShowDataTypes()
	json_handling.JsonHandling()

}
