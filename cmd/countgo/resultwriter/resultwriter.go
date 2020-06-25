package resultwriter

import (
	"fmt"
	"sync"
)

// Result represents answer received from textsearch package.
type Result struct {
	Source    string
	WordCount int
	Error     error
	EndOfData bool
}

// ResultWriterConfig contains settings for ResultWriter.
type ResultWriterConfig struct {
	// Debug ON show full error stack.
	Debug        bool
	SearchString string
}

type ResultWriter struct {
	wg sync.WaitGroup
	// chan for storing and printing result for different sources
	Results chan Result
	Config  ResultWriterConfig
}

func NewResultWriter(config ResultWriterConfig) ResultWriter {
	return ResultWriter{
		sync.WaitGroup{},
		make(chan Result),
		config,
	}
}

func (rw *ResultWriter) Close() {
	// push end indicator element
	rw.Results <- Result{"", 0, nil, true}
	close(rw.Results)
}

func (rw *ResultWriter) Run(results <-chan Result) {
	rw.wg.Add(1)
	go rw.run(results)
}

func (rw *ResultWriter) run(results <-chan Result) {
	for result := range results {
		if result.EndOfData {
			break
		}
		if result.Error != nil {
			if rw.Config.Debug {
				// show full error stack
				fmt.Printf(
					"Failed to count '%s' in %s failed: %v\n",
					rw.Config.SearchString,
					result.Source,
					result.Error,
				)
			} else {
				// show only top error
				fmt.Printf(
					"Failed to count '%s' in %s\n",
					rw.Config.SearchString,
					result.Source,
				)
			}
			continue
		}
		fmt.Printf("Count for %s: %d\n", result.Source, result.WordCount)
	}
	rw.wg.Done()
}
