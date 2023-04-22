package workerpool

import (
	"fmt"
	"time"
)

func ExampleNewWorkerPool() {
	// Define a work function that simply prints the work item
	workFunc := func(item int) {
		fmt.Println(item)
	}
	// Create a new WorkerPool with 2 workers and a channel size of 10
	wp := NewWorkerPool[int](2, 10, workFunc)
	// Submit 5 work items to the pool
	for i := 0; i < 5; i++ {
		wp.Submit(i)
	}
	time.Sleep(time.Second)
	// Shutdown the pool and wait for all workers to finish
	wp.Shutdown()

	// output:
	// 0
	// 1
	// 2
	// 3
	// 4
}
