.PHONY: docker-ci
docker-ci:
	@docker build --no-cache -t gocount-ci:0.0.0 -f docker/ci.dockerfile .

.PHONY: test
test: docker-ci
	@docker run gocount-ci:0.0.0

.PHONY: test-local
test-local:
	@go test -count=1 ./pkg/...
	@go test -count=1 ./cmd/...
	make functional-tests

.PHONY: lint
lint:
	@golangci-lint run --config .golangci.yaml

.PHONY: functional-tests
functional-tests: build
	pytest -s ./test/funcitonal_test/functional_test.py
	make clear

.PHONY: build
build:
	@go build github.com/kichyr/GoSearcher/cmd/countgo/

.PHONY: clear
clear:
	rm countgo

.PHONY: benchmark
benchmark: build
	sudo bash ./test/benchmarks/bench.sh