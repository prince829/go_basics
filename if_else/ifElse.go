package if_else

import "fmt"

func ShowIfelse() {
	var age int
	fmt.Println("Enter your age:")
	fmt.Scanln(&age)
	if age < 13 {
		fmt.Println("you are a child")
	} else if age < 18 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are a adult")
	}

}
