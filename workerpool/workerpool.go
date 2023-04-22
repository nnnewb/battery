package workerpool

import (
	"context"
	"sync"
)

// WorkerPool is a generic type representing a pool of workers that can execute a given work function on a queue of work items.
// It is parameterized by the type of the work items.
type WorkerPool[T any] struct {
	channel chan T             // channel to hold the work items
	ctx     context.Context    // context to allow for graceful shutdown
	cancel  context.CancelFunc // cancel function to cancel the context
	wg      *sync.WaitGroup    // WaitGroup to keep track of the workers
}

// NewWorkerPool creates a new WorkerPool with the given maximum number of workers, channel size, and work function.
// It starts the workers and returns a pointer to the created WorkerPool.
func NewWorkerPool[T any](workerNumMaximum int, channelSize int, workFunc func(T)) *WorkerPool[T] {
	// Create a WaitGroup to keep track of the workers
	wg := &sync.WaitGroup{}
	wg.Add(workerNumMaximum)
	// Create a context and a cancel function to allow for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	// Create a channel with the given queue size to hold the work items
	c := make(chan T, channelSize)
	// Start the workers
	for i := 0; i < workerNumMaximum; i++ {
		go func(ctx context.Context, wg *sync.WaitGroup) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case param, ok := <-c:
					if !ok {
						return
					}
					workFunc(param)
				}
			}
		}(ctx, wg)
	}

	// Return a pointer to the created WorkerPool
	return &WorkerPool[T]{
		channel: c,
		ctx:     ctx,
		cancel:  cancel,
		wg:      wg,
	}
}

// Submit adds a work item to the channel for the workers to execute.
func (t *WorkerPool[T]) Submit(item T) {
	t.channel <- item
}

// Shutdown cancels the context and waits for all workers to finish their current work items before returning.
func (t *WorkerPool[T]) Shutdown() {
	t.cancel()
	t.wg.Wait()
}
