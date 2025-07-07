package routines_concurrency

import (
	"fmt"
	"time"
)

/*
Go was built for concurrency. This means your program can do many things at once — without complex thread management like in Java or C++.

🧠 What is a Goroutine?
A goroutine is a lightweight thread managed by the Go runtime.

You create a goroutine using the go keyword:

go doSomething()
While main() is running, this doSomething() function runs in the background.
*/
func printMessage(message string) {
	for range 3 {
		fmt.Println(message)
		time.Sleep(500 * time.Millisecond)
	}
}
func RoutinesConcurrency() {
	go printMessage("From GoRoutine") //Runs on background
	printMessage("From main")         // Runs in main Thread
	time.Sleep(2 * time.Second)       // Give goroutine time to finish
}
