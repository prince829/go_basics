package functions

import "fmt"

func add(a int, b int) int {
	return a + b
}
func divide(a int, b int) (int, int) {
	quotient := a / b
	reminder := a % b
	return quotient, reminder
}

func Functions() {
	sum := add(2, 4)
	q, r := divide(7, 3)
	fmt.Println("The Total sum:", sum)
	fmt.Println("Quotient:", q, " Reminder:", r)
}
