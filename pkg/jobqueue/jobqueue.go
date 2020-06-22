package jobqueue

import (
	"sync"
)

type Job interface {
	Process()
}

type JobQueue struct {
	// use a WaitGroup for implementing waiting of ending all jobs
	wg      sync.WaitGroup
	jobChan chan (Job)
}

func NewJobQueue(workerNumber int) *JobQueue {
	queue := JobQueue{
		sync.WaitGroup{},
		make(chan Job, workerNumber),
	}
	queue.wg.Add(1)
	go queue.runWorkers()
	return &queue
}

func (jobQueue *JobQueue) PushJob(job Job) {
	jobQueue.jobChan <- job
}

//
func (jobQueue *JobQueue) Close() {
	close(jobQueue.jobChan)
	jobQueue.wg.Wait()

}

func (jobQueue *JobQueue) runWorkers() {
	for job := range jobQueue.jobChan {
		jobQueue.wg.Add(1)
		go func() {
			defer jobQueue.wg.Done()
			job.Process()
		}()
	}
	jobQueue.wg.Done()
}
