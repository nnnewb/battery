// Package workerpool provide worker pool pattern implementation.
//
// The worker pool pattern is a concurrency design pattern that involves a group of worker threads that are
// waiting for tasks to be assigned to them.
//
// The worker pool pattern is useful when there is a need to perform a large number of tasks concurrently,
// but creating a new thread for each task would be inefficient.
//
// Instead, a fixed number of worker threads are created, and tasks are assigned to them as they become available.
//
// The worker pool pattern and message queue are both used to manage and distribute workloads across multiple workers.
//
// However, there are some key differences between the two.
//
// The worker pool pattern, as implemented in the provided code, involves a fixed number of workers that are created at
// the start of the program and are responsible for executing work items that are submitted to a channel.
//
// The number of workers is determined by the workerNumMaximum parameter passed to the NewWorkerPool function.
//
// This pattern is useful when the number of workers needed is known in advance and the workload is relatively predictable.
//
// On the other hand, a message queue is a more flexible solution that allows for dynamic scaling of workers based on
// the workload.
//
// Work items are submitted to a queue, and workers can be added or removed as needed to process the items in the queue.
//
// This pattern is useful when the workload is unpredictable or varies greatly over time.
//
// In terms of similarities, both patterns involve distributing workloads across multiple workers to improve performance
// and efficiency. They also both involve some form of queueing mechanism to manage the work items.
//
// Overall, the worker pool pattern is a simpler and more rigid solution compared to a message queue, but it can be
// more efficient in certain situations where the workload is relatively predictable.
//
// A message queue, on the other hand, is a more flexible and scalable solution that can handle a wider range of workloads.
package workerpool
