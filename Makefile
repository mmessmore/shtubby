GO_FILES=$(shell find . -name '*.go' | tr '\n' ' ')
DEST:="$(HOME)"

shtubby: $(GO_FILES) .pretty
	go build -o $@

.pretty: $(GO_FILES)
	find . -name "*.go" -print0 | xargs -0 goimports -w
	go mod tidy
	touch .pretty

install: $(DEST)/bin/shtubby

$(DEST)/bin:
	mkdir -p $@

$(DEST)/bin/shtubby: $(DEST)/bin shtubby
	install -m 755 shtubby $(DEST)/bin/shtubby

.PHONY: clean
clean:
	rm -f shtubby .pretty
