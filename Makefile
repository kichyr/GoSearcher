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