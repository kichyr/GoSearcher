package jobqueue


type Job interface {
	Process() (interface, error) 
}

type JobQueue struct {
	// use a WaitGroup for implementing waiting of ending all jobs
	var wg sync.WaitGroup
	//
	jobChan chan(func)
}

func NewJobQueue struct {
	
}



func worker(jobChan <-chan Job) {
    defer wg.Done()

    for job := range jobChan {
        process(job)
    }
}

// increment the WaitGroup before starting the worker
wg.Add(1)
go worker(jobChan)

// to stop the worker, first close the job channel
close(jobChan)

// then wait using the WaitGroup
wg.Wait()