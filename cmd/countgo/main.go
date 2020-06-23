package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/kichyr/GoSearcher/pkg/jobqueue"
	"github.com/kichyr/GoSearcher/pkg/textsearch"
)

type Result struct {
	Source    string
	WordCount int
	Error     error
}

func ResultWriter(results <-chan Result) {
	for result := range results {

		if result.Error != nil {
			fmt.Printf("Failed to count 'go' in %s. %v\n", result.Source, result.Error)
			continue
		}

		fmt.Printf("Count for %s: %d\n", result.Source, result.WordCount)
	}
}

type job struct {
	source string
	result chan Result
}

func (j *job) Process() {
	count, err := textsearch.CountString("go", j.source)
	if err != nil {
		j.result <- Result{j.source, 0, err}
		return
	}
	j.result <- Result{j.source, count, nil}
}

func main() {
	var (
		workerNumber int
	)
	flag.IntVar(&workerNumber, "k", 5, "Maximum workers")
	flag.Parse()

	inputReader := bufio.NewScanner(os.Stdin)
	results := make(chan Result)

	jobs := jobqueue.NewJobQueue(workerNumber)

	go ResultWriter(results)

	for inputReader.Scan() {
		s := inputReader.Text()
		jobs.PushJob(&job{s, results})
	}

	close(results)
	jobs.Close()
}
