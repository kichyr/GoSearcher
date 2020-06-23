.PHONY: build
build:
	@go build github.com/kichyr/GoSearcher/cmd/countgo/

.PHONY: clear
clear:
	rm countgo