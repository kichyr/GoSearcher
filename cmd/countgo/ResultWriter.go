package main

import (
	"fmt"
	"sync"
)

// Result represents answer received from textsearch package
type Result struct {
	Source    string
	WordCount int
	Error     error
	endOfData bool
}

type ResultWriter struct {
	wg sync.WaitGroup
	// chan for storing and printing result for different sources
	Results chan Result
}

func NewResultWriter() ResultWriter{
	return ResultWriter{
		sync.WaitGroup{},
		make(chan Result),
	}
}

func (rw *ResultWriter) Close() {
	rw.Results <- Result{"", 0, nil, true}
	close(rw.Results)
}

func (rw *ResultWriter) Run(results <-chan Result) {
	rw.wg.Add(1)
	go rw.run(results)
}

func (rw *ResultWriter) run(results <-chan Result) {
	for result := range results {
		if result.endOfData {
			break
		}
		if result.Error != nil {
			fmt.Printf("Failed to count 'go' in %s. %v\n", result.Source, result.Error)
			continue
		}
		fmt.Printf("Count for %s: %d\n", result.Source, result.WordCount)
	}
	rw.wg.Done()
}