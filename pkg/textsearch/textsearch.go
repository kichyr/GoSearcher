package textsearch

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/*
This package provides interface for searching text in various resources
such as local files and web pages by given URL or path to local file.
*/

func CountString(searchString string, resource string) (int, error) {
	// rule to check that resource is either file path or url.
	if strings.HasPrefix(resource, "http://") || strings.HasPrefix(resource, "https://") {
		// url case
		return CountStringByURL(searchString, resource)
	} else {
		// file path case
		return CountStringByFilePath(searchString, resource)
	}
}

func CountStringByFilePath(searchString string, path string) (int, error) {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("can't open file %s: %v", path, err)
	}

	return bytes.Count(data, []byte(searchString)), nil
}

func CountStringByURL(searchString string, URL string) (int, error) {
	data, err := get(URL)
	if err != nil {
		return 0, err
	}
	return bytes.Count(data, []byte(searchString)), nil
}

func get(URL string) ([]byte, error) {
	uri, err := url.Parse(URL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL %s: %v", URL, err)
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		return nil, fmt.Errorf("http get %s failed: %v", URL, err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code is not 2**: %v, respose: %s", resp.StatusCode, respBody)
	}

	return respBody, nil
}
