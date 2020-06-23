package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/kichyr/GoSearcher/pkg/jobqueue"
	"github.com/kichyr/GoSearcher/pkg/textsearch"
)

// initial settings
const (
	searchString        = "go"
	defaultWorkerNumber = 5
)

// Implementation of the interface for working with jobqueue package.
type job struct {
	source string
	result chan Result
}

// Process wraps necessary function in method without arguments and return values
// to be executed in job queue.
func (j *job) Process() {
	count, err := textsearch.CountString(searchString, j.source)
	if err != nil {
		j.result <- Result{j.source, 0, err, false}
		return
	}
	j.result <- Result{j.source, count, nil, false}
}

func main() {
	var (
		workerNumber int
	)
	flag.IntVar(&workerNumber, "k", defaultWorkerNumber, "Maximum workers")
	flag.Parse()

	inputReader := bufio.NewScanner(os.Stdin)

	jobs := jobqueue.NewJobQueue(workerNumber)

	resWriter := NewResultWriter()

	// goroutine that prints results from results chan
	resWriter.Run(resWriter.Results)

	for inputReader.Scan() {
		s := inputReader.Text()
		// Push new job in queue.
		// If queue is full it blocks until the worker gets out.
		// If there is free worker it doesn't block and continue to read input.
		jobs.PushJob(&job{s, resWriter.Results})
	}
	fmt.Println("kek")

	jobs.Close()
	resWriter.Close()
}
