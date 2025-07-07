package channel

import "fmt"

//This is unbuffered

/*
🧠 What is a channel?
A channel lets two goroutines communicate with each other.

Think of it like a pipe where one goroutine puts data in, and another takes it out.
*/
func sayHello(ch chan string) {
	ch <- "Hello from goroutine"
}
func Channel() {
	ch := make(chan string) //Create a channel
	go func() {
		ch <- "Hello from go routine" //Send message to channel
	}()
	message := <-ch //receive message from channel
	fmt.Println(message)
	go sayHello(ch)

	message = <-ch //receive message from channel
	fmt.Println(message)
}
