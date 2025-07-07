package catch_panic_recover

import "fmt"

func safeDivision(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover:", r)
		}
	}()
	fmt.Println(a / b) //Trigger panic if b=0
}

func positiveNumber(a int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover:", r)
		}
	}()

	if a < 0 {
		panic("Can not be less than 0")
	}
	fmt.Println("Input", a)
}

func MainFunc() {
	fmt.Println("Program started:")
	positiveNumber(1)
	positiveNumber(-2) //it will panic but will be handled by recover

	safeDivision(1, 10)
	safeDivision(1, 0) //it will panic but will be handled by recover
	fmt.Println("Program Ended Safely")
}
