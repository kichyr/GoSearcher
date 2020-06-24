package jobqueue

import (
	"fmt"
	"sync"
)

/*
This package implements universal queue of workers.
Since go doesn't have generic types, to use this package
user should implement the interface Job
and wrap necessary function in Process() method.
*/

// Job describes a wrapper for user's function
type Job interface {
	Process()
}

type JobQueue struct {
	// use a WaitGroup for implementing waiting of ending all jobs
	wg      sync.WaitGroup
	jobChan chan (Job)
}

func NewJobQueue(workerNumber int) (*JobQueue, error) {
	if workerNumber < 1 {
		return nil, fmt.Errorf(
			"Give wrong worker number: %v, it should be greater than 0",
			workerNumber)
	}
	queue := JobQueue{
		sync.WaitGroup{},
		make(chan Job, workerNumber),
	}
	queue.wg.Add(1)
	go queue.runWorkers()
	return &queue, nil
}

// PushJob add new job in queue.
// It blocks when there is no free workers.
func (jobQueue *JobQueue) PushJob(job Job) {
	jobQueue.jobChan <- job
}

// Close should be called after all jobs will be pushed into queue.
// It blocks until all jobs in the queue will be done.
func (jobQueue *JobQueue) Close() {
	close(jobQueue.jobChan)
	jobQueue.wg.Wait()

}

func (jobQueue *JobQueue) runWorkers() {
	for job := range jobQueue.jobChan {
		jobQueue.wg.Add(1)
		go func(j Job) {
			defer jobQueue.wg.Done()
			j.Process()
		}(job)
	}
	jobQueue.wg.Done()
}
