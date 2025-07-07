package data_types

import "fmt"

func ShowDataTypes() {
	//string
	var name string = "Helo User"
	// int / int8, int16, int32, int64
	//Stores integers (whole numbers).
	var age int = 24 //int defaults to 32-bit or 64-bit depending on your system.

	/*
	 float32, float64
	 Stores decimal numbers.
	*/
	var height float32 = 5.09
	/*
		bool
		Store boolean value
	*/
	var isStudent bool = true
	/*
			rune and byte
		rune: represents a Unicode character (alias for int32)

		byte: represents a single byte (alias for uint8)
	*/
	var letter rune = 'A'
	var ascii byte = 'A'

	/*
		short declaration
	*/
	a := "Prince" // string
	b := 25       // int
	c := 5.9      // float64
	d := true     // bool
	// Go does not allow implicit type conversion — you must convert manually.
	var a1 int = 10
	var b1 float64 = float64(a1) // valid
	var a2 int = 10
	// var b2 float64 = a2  // ❌ invalid
	fmt.Println(letter, ascii, "assssss")
	fmt.Print(name, age, height, isStudent, letter, ascii, a, b, c, d, b1, a2)

	const ageLimit = 25
	var age1 int
	fmt.Println("Enter your age:")
	fmt.Scan(&age1)
	if age >= ageLimit {
		fmt.Println("Your are validated to access this portal")
	} else {
		fmt.Println("You had exceeded the gae limit")
	}

}
