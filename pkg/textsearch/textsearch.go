package textsearch

import "strings"

/*
This package provides interface for searching text in various resources
such as local files and web pages by given URL or path to local file.
*/

func NewTextSearcher(workerNumber int) TextSearcher {
	return TextSearcher{}
}

func CountString(resource string) (int, error) {
	// rule to check that resource is either file path or url.
	if strings.HasPrefix(resource, "http://") || strings.HasPrefix(resource, "https://") {
		// url case

	} else {
		// file path case

	}
	return
}

func countStringByFilePath() (int, error) {

}

func countStringByURL() (int, error) {

}
