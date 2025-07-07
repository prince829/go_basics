package for_loop

import "fmt"

func ForLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//To make it as while loop
	i := 5
	for i < 10 {
		fmt.Println(i)
		i++
	}
	//Infinite for loop
	for {
		fmt.Println("Running forever!")
		break // don’t forget to break eventually
	}
	//break to stop the entire loop.
	//continue to skip the current iteration and jump to the next one.
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue //it will skip even number
		}
		fmt.Println(i)
	}
}
