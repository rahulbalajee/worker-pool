package main

import (
	"fmt"
	"streamer"
)

func main() {
	// Define number of workers and jobs
	const numJobs = 4
	const numWorkers = 2

	// Create channels for work and results
	notifyChan := make(chan streamer.ProcessingMessage, numJobs)
	defer close(notifyChan)

	videoQueue := make(chan streamer.VideoProcessingJob, numJobs)
	defer close(videoQueue)

	// Get a worker pool going
	wp := streamer.New(videoQueue, numWorkers)
	fmt.Println("wp:", wp)

	// Start the worker pool
	wp.Run()

	// Create 4 videos to send to worker pool

	// Send the videos to worker pool

	// Print out results
}
