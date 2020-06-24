#  GoCount [![Build Status](https://travis-ci.org/kichyr/GoSearcher.svg?branch=master)](https://travis-ci.org/kichyr/GoSearcher)

This repo contains GoCount app that can count occurrences of string "go" on given web resources and files. Also this repo contains implementation of Job Queue package that I tried to make maximum flexible in usage and package for searching substrings on different type of resource. Both of this packages are used in GoCount app.

Also here you can find some functional and unit tests for this app and benchmark script that shows dependency between worker number and performance.


## Quick start:
```
# here k=x means number of workers
$ make build
$ echo -e 'https://golang.org/doc/effective_go.html\n./test test_files/simple_test.txt' | ./countgo -k=5
```

## Run test in docker:
```
$ make test
```

## Run tests locally:
```
$ pip3 install -r ./test/requirements.txt
$ make test-local
```

## Generate benchmark plot:
```
make benchmark
```

#Benchmark
Thats what I got on "https://golang.org/doc/effective_go.html\nhttps://ru.wikipedia.org/wiki/Go" * 200 string

![](https://raw.githubusercontent.com/kichyr/GoSearcher/master/test/benchmarks/bench.png)
