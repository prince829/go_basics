package select_function

import (
	"fmt"
	"time"
)

/*
select lets you listen to multiple channels at the same time.
It picks the first one that is ready, just like switch-case but for channels.
*/
func sendFast(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Fast channel"
}
func sendSlow(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Slow Channel"
}
func SelectFunction() {

	//Syntax
	/*
			select {
		case val := <-ch1:
			fmt.Println("Received from ch1:", val)
		case val := <-ch2:
			fmt.Println("Received from ch2:", val)
		default:
			fmt.Println("Nothing ready")
		}
	*/
	ch1 := make(chan string)
	ch2 := make(chan string)
	go sendFast(ch1)
	go sendSlow(ch2)
	select {
	case message1 := <-ch1:
		fmt.Println("Received:", message1) // it will be printed as it has less waiting time
	case message2 := <-ch2:
		fmt.Println("Received:", message2)
	}

}
