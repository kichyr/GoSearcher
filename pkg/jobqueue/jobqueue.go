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

// jobWrap add additional info to users Job interface
// implementation such as endFlag
// that becomes true when user tries to Close().
type jobWrap struct {
	job     Job
	endFlag bool
}

type JobQueue struct {
	// use a WaitGroup for implementing waiting of ending all jobs
	wg sync.WaitGroup
	// use as queue
	jobChan chan (jobWrap)
	// implementing semaphore that bound worker number
	sem chan int
}

func NewJobQueue(workerNumber int) (*JobQueue, error) {
	if workerNumber < 1 {
		return nil, fmt.Errorf(
			"Give wrong worker number: %v, it should be greater than 0",
			workerNumber)
	}
	queue := JobQueue{
		sync.WaitGroup{},
		make(chan jobWrap, workerNumber),
		make(chan int, workerNumber),
	}
	queue.wg.Add(1)
	go queue.runWorkers()
	return &queue, nil
}

// PushJob add new job in queue.
// It blocks when there is no free workers.
func (jobQueue *JobQueue) PushJob(job Job) {
	jobQueue.jobChan <- jobWrap{job, false}
}

// Close should be called after all jobs will be pushed into queue.
// It blocks until all jobs in the queue will be done.
func (jobQueue *JobQueue) Close() {
	jobQueue.jobChan <- jobWrap{nil, true}
	jobQueue.wg.Wait()
	close(jobQueue.sem)
	close(jobQueue.jobChan)
}

func (jobQueue *JobQueue) runWorkers() {
	for jobWrp := range jobQueue.jobChan {
		if jobWrp.endFlag {
			break
		}
		jobQueue.wg.Add(1)
		// acquire semaphore
		jobQueue.sem <- 1
		go func(j Job) {
			defer jobQueue.wg.Done()
			j.Process()
			// release semaphore
			<-jobQueue.sem
		}(jobWrp.job)
	}
	jobQueue.wg.Done()
}
