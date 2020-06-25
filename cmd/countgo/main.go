package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/kichyr/GoSearcher/cmd/countgo/resultwriter"
	"github.com/kichyr/GoSearcher/pkg/jobqueue"
	"github.com/kichyr/GoSearcher/pkg/textsearch"
)

// initial settings
const (
	searchString        = "go"
	defaultWorkerNumber = 5
)

// Implementation of the interface for working with jobqueue package.
type Job struct {
	source string
	result chan resultwriter.Result
}

// Process wraps necessary function in method without arguments and return values
// to be executed in job queue.
func (j Job) Process() {
	count, err := textsearch.CountString(searchString, j.source)
	if err != nil {
		j.result <- resultwriter.Result{
			Source:    j.source,
			WordCount: 0,
			Error:     err,
			EndOfData: false,
		}
		return
	}
	j.result <- resultwriter.Result{
		Source:    j.source,
		WordCount: count,
		Error:     nil,
		EndOfData: false,
	}
}

func main() {
	var (
		workerNumber int
		debug        bool
	)
	flag.IntVar(&workerNumber, "k", defaultWorkerNumber, "Maximum workers")
	flag.BoolVar(&debug, "debug", false, "Shows full error description")
	flag.Parse()
	inputReader := bufio.NewScanner(os.Stdin)

	jobs, err := jobqueue.NewJobQueue(workerNumber)
	if err != nil {
		fmt.Printf("Can't create worker queue, failed: %v", err)
		os.Exit(1)
	}

	resWriter := resultwriter.NewResultWriter(
		resultwriter.ResultWriterConfig{
			Debug:        debug,
			SearchString: searchString,
		})

	// goroutine that prints results from results chan
	resWriter.Run(resWriter.Results)

	for inputReader.Scan() {
		s := inputReader.Text()
		// Push new job in queue.
		// If queue is full it blocks until the some worker gets out.
		// If there is free worker it doesn't block
		// it starts nre job in worker and continue to read input.
		jobs.PushJob(Job{s, resWriter.Results})
	}

	jobs.Close()
	resWriter.Close()
}
