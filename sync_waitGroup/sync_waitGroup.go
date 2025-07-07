package sync_waitGroup

import (
	"fmt"
	"sync"
	"time"
)

/*
✅ What is sync.WaitGroup?
It’s like a counter that:

Increases when you start a goroutine

Decreases when a goroutine finishes

Waits until all are done ✅
*/

func printTask(name string, wg *sync.WaitGroup) {
	defer wg.Done() //Tell wait group that task has done
	for i := 0; i < 3; i++ {
		fmt.Println(name, "step ", i)
		time.Sleep(300 * time.Millisecond)
	}

}
func SyncWaitGroup() {
	//Syntax
	/*
			var wg sync.WaitGroup

		wg.Add(1)         // Register 1 goroutine
		go func() {
		    defer wg.Done() // Mark it as done
		    doSomething()
		}()

		wg.Wait()         // Block until all goroutines finish
	*/
	var wg sync.WaitGroup
	wg.Add(2) //Excepting 2 go routing
	go printTask("Task A", &wg)
	go printTask("Task B", &wg)
	wg.Wait() // Wait for both task to complete
	fmt.Println("All Task Done")
}
