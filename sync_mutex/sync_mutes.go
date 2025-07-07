package sync_mutex

import (
	"fmt"
	"sync"
)

/*
Why do we need a Mutex?
Imagine two or more goroutines modifying the same variable — this can lead to:

# Inconsistent results

# Unexpected bugs

❌ Data races
*/
var counter int
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}
func SyncMutex() {
	var wg sync.WaitGroup
	wg.Add(3)
	go increment(&wg)
	go increment(&wg)
	go increment(&wg)
	// time.Sleep(time.Second)
	wg.Wait()
	fmt.Println("Counter:", counter)
}
