package worker_pool

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
🧠 What is a Worker Pool?
It’s a pattern where:

You have a set number of workers (goroutines)

# You send them tasks through a channel

# Each worker picks up and processes a task

The main program waits until all tasks are done ✅

This is great for:

CPU-bound or IO-bound tasks (e.g. file processing, API calls)

Controlling goroutine usage to avoid overwhelming resources
*/
func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(1 * time.Second) // simulate time-consuming task
		fmt.Printf("Worker %d finished job %d\n", id, job)

	}
}

func largeWorker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		randomTime := 0.5 + rand.Float64()*(1.5-0.5)
		time.Sleep(time.Duration(randomTime * float64(time.Second)))
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}
}
func largerWorker() {
	const numJobs = 10
	const numWorker = 3
	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup
	//Start Worker
	for i := 1; i <= numWorker; i++ {
		wg.Add(1)
		go largeWorker(i, jobs, &wg)
	}
	//receive jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
	fmt.Println("Task has completed")
}
func WorkerPool() {
	const numJobs = 5
	const numWorker = 3
	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup
	//Start Worker
	for i := 1; i <= numWorker; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}
	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
	fmt.Println("Task has completed")
	largerWorker()
}
