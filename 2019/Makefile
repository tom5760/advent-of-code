BIN_DIR = bin
TOOLS_DIR = $(BIN_DIR)/tools

DAYS = $(shell find -maxdepth 1 -name "day*" | sed 's:./::')

export PATH := $(shell pwd)/$(TOOLS_DIR):$(PATH)

.PHONY: all
all:
	echo "run specific day with 'make dayXX'"

.PHONY: $(DAYS)
$(DAYS):
	go test ./$@
	@if test -f ./$@/input; then go run ./$@ < ./$@/input; else go run ./$@; fi

.PHONY: tools
tools:
	mkdir -p $(TOOLS_DIR)
	grep _ < tools.go \
		| cut --delimiter='"' --fields=2 \
		| xargs --verbose env GOBIN=$$(realpath $(TOOLS_DIR)) go install

.PHONY: test
test:
	go test ./...

.PHONY: generate
generate:
	go generate ./...
