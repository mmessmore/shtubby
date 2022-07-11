GO_FILES=$(shell find . -name '*.go' | tr '\n' ' ')

shtubby: $(GO_FILES) .pretty
	go build -o $@

.pretty: $(GO_FILES)
	find . -name "*.go" -print0 | xargs -0 goimports -w
	go mod tidy
	touch .pretty

.PHONY: clean
clean:
	rm -f schtubby .pretty
